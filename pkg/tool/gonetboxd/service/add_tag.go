package service

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (s *Service) AddTag(
	deviceName string,
	tag string,
) (*device.Device, error) {
	return s.client.AddTag(deviceName, tag)
}
