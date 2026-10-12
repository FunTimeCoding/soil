package service

import "github.com/funtimecoding/soil/pkg/netbox/constant"

func (s *Service) RemoveDeviceLabel(
	device string,
	key string,
) error {
	d, e := s.client.DeviceByName(device)

	if e != nil {
		return e
	}

	return s.store.RemoveLabel(constant.DeviceAddress, d.Identifier, key)
}
