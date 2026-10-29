package service

import "github.com/funtimecoding/soil/pkg/netbox/device_type"

func (s *Service) DeviceTypes() ([]*device_type.Type, error) {
	return s.client.DeviceTypes()
}
