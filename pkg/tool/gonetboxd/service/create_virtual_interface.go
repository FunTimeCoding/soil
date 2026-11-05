package service

import "github.com/netbox-community/go-netbox/v4"

func (s *Service) CreateVirtualInterface(
	machine string,
	name string,
) (*netbox.VMInterface, error) {
	return guarded(
		s.guard,
		func() (*netbox.VMInterface, error) {
			m, e := s.client.VirtualMachineByName(machine)

			if e != nil {
				return nil, e
			}

			return s.client.CreateVirtualInterface(m, name)
		},
	)
}
