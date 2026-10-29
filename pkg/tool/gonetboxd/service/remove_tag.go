package service

import "github.com/funtimecoding/soil/pkg/netbox/device"

func (s *Service) RemoveTag(
	deviceName string,
	tag string,
) (*device.Device, error) {
	return s.client.RemoveTag(deviceName, tag)
}
