package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) PhysicalAddresses() ([]*physical_address.Address, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*physical_address.Address, error) {
			return s.client.PhysicalAddresses()
		},
	)
}
