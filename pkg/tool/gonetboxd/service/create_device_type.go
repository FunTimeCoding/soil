package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateDeviceType(
	model string,
	manufacturer string,
) (*device_type.Type, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device_type.Type, error) {
			m, e := s.client.Manufacturer(manufacturer)

			if e != nil {
				return nil, e
			}

			return s.client.CreateDeviceType(model, m)
		},
	)
}
