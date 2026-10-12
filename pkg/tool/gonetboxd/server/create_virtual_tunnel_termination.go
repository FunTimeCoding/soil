package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateVirtualTunnelTermination(
	_ context.Context,
	r server.CreateVirtualTunnelTerminationRequestObject,
) (server.CreateVirtualTunnelTerminationResponseObject, error) {
	t, e := s.service.CreateVirtualTunnelTermination(
		r.Name,
		r.Body.Tunnel,
		r.Body.Interface,
		r.Body.Role,
	)

	if e != nil {
		return server.CreateVirtualTunnelTermination500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.CreateVirtualTunnelTermination201JSONResponse(
		*convert.TunnelTermination(t),
	), nil
}
