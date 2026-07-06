package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateDeviceType(
	_ context.Context,
	r server.CreateDeviceTypeRequestObject,
) (server.CreateDeviceTypeResponseObject, error) {
	m, e := s.client.ManufacturerByName(r.Body.Manufacturer)

	if e != nil {
		return server.CreateDeviceType500JSONResponse(*s.captureDetail(e)), nil
	}

	t, f := s.client.CreateDeviceType(r.Body.Model, m)

	if f != nil {
		return server.CreateDeviceType500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateDeviceType201JSONResponse{
		Identifier:   t.Identifier,
		Model:        t.Model,
		Manufacturer: &m.Name,
	}, nil
}
