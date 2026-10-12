package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_disk"

func (s *Service) VirtualDisks(machine string) ([]*virtual_disk.Disk, error) {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.client.VirtualMachineDisks(m)
}
