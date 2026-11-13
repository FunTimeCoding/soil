package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateSite(name string) (*site.Site, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*site.Site, error) {
			return s.client.CreateSite(name)
		},
	)
}
