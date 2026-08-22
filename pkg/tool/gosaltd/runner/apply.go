package runner

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	provision "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"slices"
	"strings"
	"time"
)

func (r *Runner) apply(
	parameters map[string]any,
	triggerSource string,
) any {
	head := r.provision.CurrentHead()
	target := "*"

	if v, okay := parameters[constant.Target]; okay {
		target = v.(string)
	}

	record := r.store.NewRun()
	record.Scope = target
	record.TriggerSource = triggerSource
	record.Status = provision.StoreStatusRunning
	record.GitHead = head
	r.store.Create(record)
	r.logger.Structured("highstate_start", constant.Target, target)
	start := time.Now()
	result, e := r.salt.Highstate(target)
	record.DurationMillisecond = time.Since(start).Milliseconds()

	if e != nil {
		record.Status = provision.StoreStatusError
		record.ErrorOutput = e.Error()
		r.logger.Structured("highstate_error", "error", e.Error())
	} else {
		output, marshalError := json.MarshalIndent(result, "", "  ")
		errors.PanicOnError(marshalError)
		record.Output = string(output)
		var silent []string

		for minion, v := range result {
			if !v.Responded {
				silent = append(silent, minion)
			}
		}

		if len(silent) > 0 {
			slices.Sort(silent)
			record.Status = provision.StoreStatusError
			record.ErrorOutput = fmt.Sprintf(
				"minion did not respond: %s",
				strings.Join(silent, ", "),
			)
			r.logger.Structured(
				"highstate_silent_minion",
				constant.Minions,
				strings.Join(silent, ", "),
			)
		} else {
			record.Status = provision.StoreStatusSuccess
			r.logger.Structured("highstate_done")
		}
	}

	r.store.Update(record)

	return record
}
