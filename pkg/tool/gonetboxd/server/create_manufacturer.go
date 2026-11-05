package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateManufacturer(
	x context.Context,
	r server.CreateManufacturerRequestObject,
) (server.CreateManufacturerResponseObject, error) {
	m, e := s.service.CreateManufacturer(r.Body.Name)

	if validation.Is(e) {
		return server.CreateManufacturer400JSONResponse(*clientError(e)),
			nil
	}

	if unavailable.Is(e) {
		return server.CreateManufacturer503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateManufacturer500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateManufacturer201JSONResponse{
		Identifier: m.Identifier,
		Name:       m.Name,
	}, nil
}
