package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (s *Service) CreateVirtualAddress(
	machine string,
	interfaceName string,
	address string,
	status string,
) (*internet_address.Address, error) {
	m, e := s.client.VirtualMachineByName(machine)

	if e != nil {
		return nil, e
	}

	i, e := s.client.VirtualMachineInterfaceByName(m, interfaceName)

	if e != nil {
		return nil, e
	}

	return s.client.CreateVirtualAddress(i.GetId(), address, status)
}
