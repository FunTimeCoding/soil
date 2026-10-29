package service

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (s *Service) SetDeviceCustomFields(
	name string,
	fields map[string]any,
) (*device.Device, error) {
	return s.client.SetDeviceCustomFields(name, fields)
}
