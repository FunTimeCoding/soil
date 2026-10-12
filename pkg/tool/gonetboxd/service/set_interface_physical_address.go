package service

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (s *Service) SetInterfacePhysicalAddress(
	device string,
	interfaceName string,
	address string,
) (*network.Interface, error) {
	h, e := physicalAddress(address)

	if e != nil {
		return nil, e
	}

	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	i, e := s.client.DeviceInterfaceByName(d, interfaceName)

	if e != nil {
		return nil, e
	}

	return s.client.UpdateInterface(d, interfaceName, i.Type, h)
}
