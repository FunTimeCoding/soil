package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
)

func (s *Service) CreateDevice(
	name string,
	role string,
	deviceType string,
	site string,
	tenantName string,
	tags []string,
) (*device.Device, error) {
	r, e := s.client.DeviceRoleByName(role)

	if e != nil {
		return nil, e
	}

	t, e := s.client.DeviceTypeByName(deviceType)

	if e != nil {
		return nil, e
	}

	i, e := s.client.SiteByName(site)

	if e != nil {
		return nil, e
	}

	var n *tenant.Tenant

	if tenantName != "" {
		n, e = s.client.TenantByName(tenantName)

		if e != nil {
			return nil, e
		}
	}

	return s.client.CreateDevice(name, r, tags, t, i, n)
}
