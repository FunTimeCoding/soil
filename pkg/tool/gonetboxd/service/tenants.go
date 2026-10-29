package service

import "github.com/funtimecoding/soil/pkg/netbox/tenant"

func (s *Service) Tenants() ([]*tenant.Tenant, error) {
	return s.client.Tenants()
}
