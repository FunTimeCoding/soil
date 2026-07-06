package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/luthermonson/go-proxmox"
)

func (s *Service) CloneMachine(
	c face.ProxmoxClient,
	identifier int,
	node string,
	options *proxmox.VirtualMachineCloneOptions,
) (int, error) {
	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return 0, e
	}

	newID, task, e := c.CloneMachine(vm, options)

	if e != nil {
		return 0, e
	}

	e = c.WaitForTask(task, 600)

	if e != nil {
		return 0, e
	}

	return newID, nil
}
