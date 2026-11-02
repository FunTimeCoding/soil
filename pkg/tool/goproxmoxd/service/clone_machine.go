package service

import (
	"github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) CloneMachine(
	instance string,
	identifier int,
	node string,
	options *proxmox.VirtualMachineCloneOptions,
) (int, error) {
	return guarded(
		s,
		instance,
		func(c face.ProxmoxClient) (int, error) {
			vm, e := findMachine(c, identifier, node)

			if e != nil {
				return 0, e
			}

			newIdentifier, task, e := c.CloneMachine(vm, options)

			if e != nil {
				return 0, e
			}

			if f := c.WaitForTask(task, 600); f != nil {
				return 0, f
			}

			return newIdentifier, nil
		},
	)
}
