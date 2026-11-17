package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePowerOutlet(
	x context.Context,
	r server.CreatePowerOutletRequestObject,
) (server.CreatePowerOutletResponseObject, error) {
	o, e := s.service.CreatePowerOutlet(r.Name, r.Body.Name)

	if not_found.Is(e) {
		return server.CreatePowerOutlet404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreatePowerOutlet400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreatePowerOutlet503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreatePowerOutlet500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreatePowerOutlet201JSONResponse(*convert.PowerOutlet(o)), nil
}
