package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDevice(
	x context.Context,
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

	if validation.Is(e) {
		return server.CreateDevice400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateDevice503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateDevice500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateDevice201JSONResponse(*convert.Device(d)), nil
}
