package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) DeviceLabels(device string) ([]*label.Label, error) {
	d, e := s.client.DeviceByName(device)

	if e != nil {
		return nil, e
	}

	return s.store.Labels(constant.DeviceAddress, d.Identifier)
}
