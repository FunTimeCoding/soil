package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateAddress(
	device string,
	interfaceName string,
	address string,
	status string,
) (*internet_address.Address, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*internet_address.Address, error) {
			d, e := s.client.Device(device)

			if e != nil {
				return nil, e
			}

			i, e := s.client.DeviceInterface(d, interfaceName)

			if e != nil {
				return nil, e
			}

			return s.client.CreateAddress(i.Identifier, address, status)
		},
	)
}
