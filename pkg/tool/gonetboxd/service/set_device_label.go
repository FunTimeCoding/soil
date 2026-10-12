package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) SetDeviceLabel(
	device string,
	key string,
	value string,
) (*label.Label, error) {
	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	return s.store.SetLabel(constant.DeviceAddress, d.Identifier, key, value)
}
