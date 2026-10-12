package service

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (s *Service) VirtualInterfaces(machine string) ([]*network.Interface, error) {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	return s.client.VirtualMachineInterfaces(m)
}
