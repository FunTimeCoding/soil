package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DeviceInterfaces(device string) ([]*network.Interface, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*network.Interface, error) {
			d, e := s.client.Device(device)

			if e != nil {
				return nil, e
			}

			return s.client.DeviceInterfaces(d.Identifier)
		},
	)
}
