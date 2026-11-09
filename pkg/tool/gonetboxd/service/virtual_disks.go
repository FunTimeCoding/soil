package service

import "github.com/funtimecoding/soil/pkg/netbox/virtual_disk"

func (s *Service) VirtualDisks(machine string) ([]*virtual_disk.Disk, error) {
	return guarded(
		s.guard,
		func() ([]*virtual_disk.Disk, error) {
			m, e := s.client.VirtualMachine(machine)

			if e != nil {
				return nil, e
			}

			return s.client.VirtualMachineDisks(m)
		},
	)
}
