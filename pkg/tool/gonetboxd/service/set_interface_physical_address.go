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

	return guarded(
		s.guard,
		func() (*network.Interface, error) {
			d, f := s.client.DeviceByName(device)

			if f != nil {
				return nil, f
			}

			i, f := s.client.DeviceInterfaceByName(d, interfaceName)

			if f != nil {
				return nil, f
			}

			return s.client.UpdateInterface(d, interfaceName, i.Type, h)
		},
	)
}
