package runner

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
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
	record.Status = store.StatusRunning
	record.GitHead = head
	r.store.Create(record)
	r.logger.Structured("highstate_start", constant.Target, target)
	start := time.Now()
	result, e := r.salt.Highstate(target)
	record.DurationMillisecond = time.Since(start).Milliseconds()

	if e != nil {
		record.Status = store.StatusError
		record.ErrorOutput = e.Error()
		r.logger.Structured(
			"highstate_error",
			"error",
			e.Error(),
		)
	} else {
		output, marshalError := json.MarshalIndent(result, "", "  ")
		errors.PanicOnError(marshalError)
		record.Output = string(output)
		record.Status = store.StatusSuccess
		r.logger.Structured("highstate_done")
	}

	r.store.Update(record)

	return record
}
