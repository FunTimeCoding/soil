package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/power_port"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreatePowerPort(
	device string,
	name string,
) (*power_port.Port, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*power_port.Port, error) {
			d, e := s.client.Device(device)

			if e != nil {
				return nil, e
			}

			return s.client.CreatePowerPort(d, name)
		},
	)
}
