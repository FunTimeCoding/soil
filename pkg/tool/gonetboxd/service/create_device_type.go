package service

import "github.com/funtimecoding/soil/pkg/netbox/device_type"

func (s *Service) CreateDeviceType(
	model string,
	manufacturer string,
) (*device_type.Type, error) {
	m, e := s.client.ManufacturerByName(manufacturer)

	if e != nil {
		return nil, e
	}

	return s.client.CreateDeviceType(model, m)
}
