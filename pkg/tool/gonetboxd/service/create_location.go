package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/location"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateLocation(
	name string,
	siteName string,
) (*location.Location, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*location.Location, error) {
			return s.client.CreateLocation(name, siteName)
		},
	)
}
