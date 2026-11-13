package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) Locations() ([]*location.Location, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*location.Location, error) {
			return s.client.Locations()
		},
	)
}
