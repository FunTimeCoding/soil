package service

import "github.com/funtimecoding/soil/pkg/netbox/network"

func (s *Service) DeviceInterfaces(device string) ([]*network.Interface, error) {
	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	return s.client.DeviceInterfaces(d.Identifier)
}
