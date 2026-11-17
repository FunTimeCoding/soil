package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cable"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreatePowerCable(
	outletDevice string,
	outlet string,
	portDevice string,
	port string,
) (*cable.Cable, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*cable.Cable, error) {
			a, e := s.client.Device(outletDevice)

			if e != nil {
				return nil, e
			}

			o, e := s.client.DevicePowerOutlet(a, outlet)

			if e != nil {
				return nil, e
			}

			b, e := s.client.Device(portDevice)

			if e != nil {
				return nil, e
			}

			p, e := s.client.DevicePowerPort(b, port)

			if e != nil {
				return nil, e
			}

			return s.client.CreatePowerCable(o, p)
		},
	)
}
