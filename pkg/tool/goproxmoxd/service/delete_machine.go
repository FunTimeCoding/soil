package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
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
		return conflict.Format(
			"machine %d is running - stop it before deleting",
			identifier,
		)
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
