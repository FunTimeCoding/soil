package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) CloneMachine(
	instance string,
	identifier int,
	node string,
	options *proxmox.VirtualMachineCloneOptions,
) (int, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return 0, clientFail
	}

	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return 0, e
	}

	newIdentifier, task, e := c.CloneMachine(vm, options)

	if e != nil {
		return 0, e
	}

	e = c.WaitForTask(task, 600)

	if e != nil {
		return 0, e
	}

	return newIdentifier, nil
}
