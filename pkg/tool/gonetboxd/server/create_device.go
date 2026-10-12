package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDevice(
	_ context.Context,
	r server.CreateDeviceRequestObject,
) (server.CreateDeviceResponseObject, error) {
	var tags []string

	if r.Body.Tags != nil {
		tags = *r.Body.Tags
	}

	tenant := ""

	if r.Body.Tenant != nil {
		tenant = *r.Body.Tenant
	}

	d, e := s.service.CreateDevice(
		r.Body.Name,
		r.Body.Role,
		r.Body.Type,
		r.Body.Site,
		tenant,
		tags,
	)

	if e != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateDevice201JSONResponse(*convert.Device(d)), nil
}
