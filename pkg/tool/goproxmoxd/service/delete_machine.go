package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) DeleteMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
	purge bool,
) error {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return e
	}

	if vm.Status == "running" {
		return constant.ErrorMachineRunning
	}

	task, e := c.DeleteMachine(
		vm,
		&proxmox.VirtualMachineDeleteOptions{
			Purge:                    proxmox.IntOrBool(purge),
			DestroyUnreferencedDisks: proxmox.IntOrBool(true),
		},
	)

	if e != nil {
		return e
	}

	return c.WaitForTask(task, 120)
}
