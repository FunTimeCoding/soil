package service

import "github.com/funtimecoding/soil/pkg/netbox/tenant"

func (s *Service) CreateTenant(name string) (*tenant.Tenant, error) {
	return s.client.CreateTenant(name)
}
