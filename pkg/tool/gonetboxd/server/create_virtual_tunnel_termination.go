package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualTunnelTermination(
	x context.Context,
	r server.CreateVirtualTunnelTerminationRequestObject,
) (server.CreateVirtualTunnelTerminationResponseObject, error) {
	t, e := s.service.CreateVirtualTunnelTermination(
		r.Name,
		r.Body.Tunnel,
		r.Body.Interface,
		r.Body.Role,
	)

	if not_found.Is(e) {
		return server.CreateVirtualTunnelTermination404JSONResponse(
				*clientError(e),
			),
			nil
	}

	if validation.Is(e) {
		return server.CreateVirtualTunnelTermination400JSONResponse(
			*clientError(e),
		), nil
	}

	if unavailable.Is(e) {
		return server.CreateVirtualTunnelTermination503JSONResponse{
			Error: e.Error(),
		}, nil
	}

	if e != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(
				*s.captureDetail(x, e),
			),
			nil
	}

	return server.CreateVirtualTunnelTermination201JSONResponse(
			*convert.TunnelTermination(t),
		),
		nil
}
