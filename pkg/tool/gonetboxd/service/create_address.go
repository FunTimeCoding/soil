package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (s *Service) CreateAddress(
	device string,
	interfaceName string,
	address string,
	status string,
) (*internet_address.Address, error) {
	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	i, e := s.client.DeviceInterfaceByName(d, interfaceName)

	if e != nil {
		return nil, e
	}

	return s.client.CreateAddress(i.Identifier, address, status)
}
