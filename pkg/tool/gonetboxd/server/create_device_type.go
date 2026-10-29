package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceType(
	_ context.Context,
	r server.CreateDeviceTypeRequestObject,
) (server.CreateDeviceTypeResponseObject, error) {
	t, e := s.service.CreateDeviceType(r.Body.Model, r.Body.Manufacturer)

	if validation.Is(e) {
		return server.CreateDeviceType400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateDeviceType403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateDeviceType500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateDeviceType201JSONResponse(*convert.DeviceType(t)), nil
}
