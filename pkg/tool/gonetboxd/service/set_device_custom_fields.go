package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SetDeviceCustomFields(
	name string,
	fields map[string]any,
) (*device.Device, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device.Device, error) {
			return s.client.SetDeviceCustomFields(name, fields)
		},
	)
}
