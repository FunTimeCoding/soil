package runner

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
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

	inventoryDirectory, e := r.renderInventory()

	if e != nil {
		return r.failInventory(e, triggerSource, head)
	}

	if inventoryDirectory != "" {
		defer system.Remove(inventoryDirectory)
	}

	var results []*run.Run

	for _, p := range playbooks {
		record := r.store.NewRun()
		record.Scope = p
		record.TriggerSource = triggerSource
		record.Status = provisionConstant.StoreStatusRunning
		record.GitHead = head
		r.store.Create(record)
		r.logger.Structured("playbook_start", constant.Playbook, p)
		start := time.Now()
		c := r.newRun().NoPanic()
		c.Directory = directory
		c.Start(command(inventoryDirectory, p)...)
		record.DurationMillisecond = time.Since(start).Milliseconds()
		record.Output = c.OutputString
		record.ErrorOutput = c.ErrorString

		if c.Error != nil {
			record.Status = provisionConstant.StoreStatusError
			r.logger.Structured(
				"playbook_error",
				constant.Playbook,
				p,
				"error",
				c.Error.Error(),
			)
		} else {
			record.Status = provisionConstant.StoreStatusSuccess
			r.logger.Structured("playbook_done", constant.Playbook, p)
		}

		r.store.Update(record)
		results = append(results, record)
	}

	return results
}
