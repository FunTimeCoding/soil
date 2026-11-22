package runner

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
	"strings"
	"time"
)

func (r *Runner) hostCommandRun(
	command string,
	host string,
	source string,
) (*run.Run, error) {
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
	pattern := fmt.Sprintf(constant.HostCommandPattern, d.Group, host)
	list := r.newRun().NoPanic()
	list.Directory = directory
	list.Start(adHoc(inventoryDirectory, pattern, constant.ListHostsFlag)...)

	if list.Error != nil {
		return nil, fmt.Errorf(
			constant.ListHostsFailed,
			list.Error,
			strings.TrimSpace(list.ErrorString),
		)
	}

	if !listed(list.OutputString, host) {
		return nil, validation.New(constant.HostCommandOutside, host, d.Group)
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

	return record, nil
}
