package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTunnel(
	x context.Context,
	r server.CreateTunnelRequestObject,
) (server.CreateTunnelResponseObject, error) {
	t, e := s.service.CreateTunnel(
		r.Body.Name,
		r.Body.Encapsulation,
		r.Body.Group,
	)

	if validation.Is(e) {
		return server.CreateTunnel400JSONResponse(*clientError(e)), nil
	}

	if unavailable.Is(e) {
		return server.CreateTunnel503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CreateTunnel500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.CreateTunnel201JSONResponse(*convert.Tunnel(t)), nil
}
