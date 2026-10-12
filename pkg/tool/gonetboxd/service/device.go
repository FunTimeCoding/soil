package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Service) Device(name string) (*device.Device, []*label.Label, error) {
	d, e := s.client.DeviceByName(name)

	if e != nil {
		return nil, nil, e
	}

	labels, e := s.store.Labels(constant.DeviceAddress, d.Identifier)

	if e != nil {
		return nil, nil, e
	}

	return d, labels, nil
}
