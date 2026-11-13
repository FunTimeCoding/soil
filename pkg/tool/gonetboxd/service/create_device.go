package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateDevice(
	name string,
	role string,
	deviceType string,
	site string,
	tenantName string,
	tags []string,
) (*device.Device, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*device.Device, error) {
			r, e := s.client.DeviceRole(role)

			if e != nil {
				return nil, e
			}

			t, e := s.client.DeviceType(deviceType)

			if e != nil {
				return nil, e
			}

			i, e := s.client.Site(site)

			if e != nil {
				return nil, e
			}

			var n *tenant.Tenant

			if tenantName != "" {
				n, e = s.client.Tenant(tenantName)

				if e != nil {
					return nil, e
				}
			}

			return s.client.CreateDevice(name, r, tags, t, i, n)
		},
	)
}
