package service

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (s *Service) CreateCable(
	deviceA string,
	interfaceA string,
	deviceB string,
	interfaceB string,
) (*cable.Cable, error) {
	a, e := s.client.DeviceByName(deviceA)

	if e != nil {
		return nil, e
	}

	sideA, e := s.client.DeviceInterfaceByName(a, interfaceA)

	if e != nil {
		return nil, e
	}

	b, e := s.client.DeviceByName(deviceB)

	if e != nil {
		return nil, e
	}

	sideB, e := s.client.DeviceInterfaceByName(b, interfaceB)

	if e != nil {
		return nil, e
	}

	return s.client.CreateCable(sideA, sideB)
}
