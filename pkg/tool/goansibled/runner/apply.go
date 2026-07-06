package runner

import (
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/goansibled/constant"
	"path/filepath"
	"time"
)

func (r *Runner) apply(
	parameters map[string]any,
	triggerSource string,
) any {
	directory := filepath.Join(r.clonePath, r.ansiblePath)
	head := r.provision.CurrentHead()
	playbooks := r.playbook

	if v, okay := parameters[constant.Playbook]; okay {
		playbooks = []string{v.(string)}
	}

	var results []*store.Run

	for _, p := range playbooks {
		record := r.store.NewRun()
		record.Scope = p
		record.TriggerSource = triggerSource
		record.Status = store.StatusRunning
		record.GitHead = head
		r.store.Create(record)
		r.logger.Structured("playbook_start", constant.Playbook, p)
		start := time.Now()
		c := r.newRun().NoPanic()
		c.Directory = directory
		c.Start("ansible-playbook", p)
		record.DurationMillisecond = time.Since(start).Milliseconds()
		record.Output = c.OutputString
		record.ErrorOutput = c.ErrorString

		if c.Error != nil {
			record.Status = store.StatusError
			r.logger.Structured(
				"playbook_error",
				constant.Playbook,
				p,
				"error",
				c.Error.Error(),
			)
		} else {
			record.Status = store.StatusSuccess
			r.logger.Structured("playbook_done", constant.Playbook, p)
		}

		r.store.Update(record)
		results = append(results, record)
	}

	return results
}
