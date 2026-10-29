package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) RemoveDeviceTag(
	_ context.Context,
	r server.RemoveDeviceTagRequestObject,
) (server.RemoveDeviceTagResponseObject, error) {
	d, e := s.service.RemoveTag(r.Name, r.Tag)

	if not_found.Is(e) {
		return server.RemoveDeviceTag404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.RemoveDeviceTag400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.RemoveDeviceTag403JSONResponse(*clientError(e)), nil
	}

	if conflict.Is(e) {
		return server.RemoveDeviceTag409JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.RemoveDeviceTag500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.RemoveDeviceTag200JSONResponse(*convert.Device(d)), nil
}
