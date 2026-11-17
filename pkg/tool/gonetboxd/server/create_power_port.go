package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePowerPort(
	x context.Context,
	r server.CreatePowerPortRequestObject,
) (server.CreatePowerPortResponseObject, error) {
	p, e := s.service.CreatePowerPort(r.Name, r.Body.Name)

	if not_found.Is(e) {
		return server.CreatePowerPort404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreatePowerPort400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreatePowerPort503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreatePowerPort500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreatePowerPort201JSONResponse(*convert.PowerPort(p)), nil
}
