package service

import "github.com/luthermonson/go-proxmox"

func (s *Service) ListMachineSnapshots(
	instance string,
	identifier int,
	node string,
) ([]*proxmox.VirtualMachineSnapshot, error) {
	c, clientFail := s.Client(instance)

	if clientFail != nil {
		return nil, clientFail
	}

	vm, e := findMachine(c, identifier, node)

	if e != nil {
		return nil, e
	}

	return c.MachineSnapshots(vm)
}
