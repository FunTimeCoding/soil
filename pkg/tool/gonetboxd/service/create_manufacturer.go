package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateManufacturer(name string) (*manufacturer.Manufacturer, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*manufacturer.Manufacturer, error) {
			return s.client.CreateManufacturer(name)
		},
	)
}
