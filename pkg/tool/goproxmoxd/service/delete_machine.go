package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) DeleteMachine(
	instance string,
	identifier int,
	node string,
	purge bool,
) error {
	return s.attempt(
		instance,
		func(c face.ProxmoxClient) error {
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
		},
	)
}
