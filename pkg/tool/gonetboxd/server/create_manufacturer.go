package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateManufacturer(
	_ context.Context,
	r server.CreateManufacturerRequestObject,
) (server.CreateManufacturerResponseObject, error) {
	m, e := s.client.CreateManufacturer(r.Body.Name)

	if e != nil {
		return server.CreateManufacturer500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateManufacturer201JSONResponse{
		Identifier: m.Identifier, Name: m.Name,
	}, nil
}
