package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceType(
	_ context.Context,
	r server.CreateDeviceTypeRequestObject,
) (server.CreateDeviceTypeResponseObject, error) {
	t, e := s.service.CreateDeviceType(r.Body.Model, r.Body.Manufacturer)

	if e != nil {
		return server.CreateDeviceType500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateDeviceType201JSONResponse(*convert.DeviceType(t)), nil
}
