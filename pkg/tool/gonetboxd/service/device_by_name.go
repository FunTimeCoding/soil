package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) deviceByName(name string) (*device.Device, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device.Device, error) {
			return s.client.Device(name)
		},
	)
}
