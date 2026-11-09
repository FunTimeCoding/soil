package service

import "github.com/funtimecoding/soil/pkg/netbox/cable"

func (s *Service) CreateCable(
	deviceA string,
	interfaceA string,
	deviceB string,
	interfaceB string,
) (*cable.Cable, error) {
	return guarded(
		s.guard,
		func() (*cable.Cable, error) {
			a, e := s.client.Device(deviceA)

			if e != nil {
				return nil, e
			}

			sideA, e := s.client.DeviceInterface(a, interfaceA)

			if e != nil {
				return nil, e
			}

			b, e := s.client.Device(deviceB)

			if e != nil {
				return nil, e
			}

			sideB, e := s.client.DeviceInterface(b, interfaceB)

			if e != nil {
				return nil, e
			}

			return s.client.CreateCable(sideA, sideB)
		},
	)
}
