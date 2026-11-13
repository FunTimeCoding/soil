package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) DeviceAddresses(device string) ([]*internet_address.Address, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*internet_address.Address, error) {
			return s.client.DeviceAddresses(device)
		},
	)
}
