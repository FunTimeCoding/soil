package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateTenant(name string) (*tenant.Tenant, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tenant.Tenant, error) {
			return s.client.CreateTenant(name)
		},
	)
}
