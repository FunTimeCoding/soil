package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/tenant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDevice(
	_ context.Context,
	r server.CreateDeviceRequestObject,
) (server.CreateDeviceResponseObject, error) {
	role, e := s.client.DeviceRoleByName(r.Body.Role)

	if e != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(e)), nil
	}

	deviceType, f := s.client.DeviceTypeByName(r.Body.Type)

	if f != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(f)), nil
	}

	site, g := s.client.SiteByName(r.Body.Site)

	if g != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(g)), nil
	}

	var tags []string

	if r.Body.Tags != nil {
		tags = *r.Body.Tags
	}

	var ten *tenant.Tenant

	if r.Body.Tenant != nil && *r.Body.Tenant != "" {
		t, h := s.client.TenantByName(*r.Body.Tenant)

		if h != nil {
			return server.CreateDevice500JSONResponse(*s.captureDetail(h)), nil
		}

		ten = t
	}

	d, i := s.client.CreateDevice(
		r.Body.Name,
		role,
		tags,
		deviceType,
		site,
		ten,
	)

	if i != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(i)), nil
	}

	return server.CreateDevice201JSONResponse(*convert.Device(d)), nil
}
