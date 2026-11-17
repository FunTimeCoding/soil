package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/power_outlet"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DevicePowerOutlets(
	device string,
) ([]*power_outlet.Outlet, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*power_outlet.Outlet, error) {
			d, e := s.client.Device(device)

			if e != nil {
				return nil, e
			}

			return s.client.DevicePowerOutlets(d.Identifier)
		},
	)
}
