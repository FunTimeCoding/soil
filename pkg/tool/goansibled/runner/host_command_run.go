package runner

import (
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/command_run"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
	"time"
)

func (r *Runner) hostCommandRun(
	command string,
	host string,
	source string,
) (*command_run.Run, error) {
	d, e := r.definition(command)

	if e != nil {
		return nil, e
	}

	inventoryDirectory, f := r.renderInventory()

	if f != nil {
		return nil, f
	}

	if inventoryDirectory != "" {
		defer system.Remove(inventoryDirectory)
	}

	directory := filepath.Join(r.cloneDirectory, r.ansibleDirectory)
	pattern, g := r.hostPattern(d, host, directory, inventoryDirectory)

	if g != nil {
		return nil, g
	}

	record := r.store.NewRun()
	record.Scope = command
	record.TriggerSource = source
	record.Status = provisionConstant.StoreStatusRunning
	record.GitHead = r.provision.CurrentHead()
	r.store.Create(record)
	start := time.Now()
	c := r.newRun().NoPanic()
	c.Directory = directory
	c.Environment(constant.CallbackEnvironment, constant.NotationCallback)
	c.Environment(constant.LoadCallbackEnvironment, constant.LoadCallback)
	c.Start(
		adHoc(
			inventoryDirectory,
			pattern,
			constant.ModuleFlag,
			d.Module,
			constant.ArgumentFlag,
			d.Arguments,
		)...,
	)
	record.DurationMillisecond = time.Since(start).Milliseconds()
	record.Output = c.OutputString
	record.ErrorOutput = c.ErrorString
	record.Status = provisionConstant.StoreStatusSuccess

	if c.Error != nil {
		record.Status = provisionConstant.StoreStatusError
	}

	r.store.Update(record)

	return command_run.New(record, report(c.OutputString)), nil
}
