package runner

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/constant"
	"path/filepath"
	"time"
)

func (r *Runner) apply(
	parameters map[string]any,
	triggerSource string,
) any {
	directory := filepath.Join(r.clonePath, r.terraformPath)
	head := r.provision.CurrentHead()
	record := r.store.NewRun()
	record.TriggerSource = triggerSource
	record.Status = store.StatusRunning
	record.GitHead = head
	r.store.Create(record)
	r.logger.Structured("terraform_apply_start")
	arguments := []string{"terraform", "apply", "-auto-approve"}

	if v, okay := parameters[constant.Target]; okay {
		record.Scope = v.(string)
		arguments = append(arguments, fmt.Sprintf("-target=%s", v.(string)))
	}

	start := time.Now()
	c := r.newRun().NoPanic()
	c.Directory = directory
	c.Start(arguments...)
	record.DurationMillisecond = time.Since(start).Milliseconds()
	record.Output = c.OutputString
	record.ErrorOutput = c.ErrorString

	if c.Error != nil {
		record.Status = store.StatusError
		r.logger.Structured(
			"terraform_apply_error",
			"error",
			c.Error.Error(),
		)
	} else {
		record.Status = store.StatusSuccess
		r.logger.Structured("terraform_apply_done")
	}

	r.store.Update(record)

	return record
}
