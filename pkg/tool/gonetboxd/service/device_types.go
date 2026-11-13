package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DeviceTypes() ([]*device_type.Type, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*device_type.Type, error) {
			return s.client.DeviceTypes()
		},
	)
}
