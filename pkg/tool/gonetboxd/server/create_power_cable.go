package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePowerCable(
	x context.Context,
	r server.CreatePowerCableRequestObject,
) (server.CreatePowerCableResponseObject, error) {
	result, e := s.service.CreatePowerCable(
		r.Body.OutletDevice,
		r.Body.Outlet,
		r.Body.PortDevice,
		r.Body.Port,
	)

	if not_found.Is(e) {
		return server.CreatePowerCable404JSONResponse(*clientError(e)), nil
	}

	if validation.Is(e) {
		return server.CreatePowerCable400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreatePowerCable503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreatePowerCable500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.CreatePowerCable201JSONResponse(*convert.Cable(result)), nil
}
