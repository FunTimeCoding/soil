package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTunnel(
	_ context.Context,
	r server.CreateTunnelRequestObject,
) (server.CreateTunnelResponseObject, error) {
	g, e := s.client.TunnelGroupByName(r.Body.Group)

	if e != nil {
		return server.CreateTunnel500JSONResponse(*s.captureDetail(e)), nil
	}

	t, f := s.client.CreateTunnel(r.Body.Name, r.Body.Encapsulation, g)

	if f != nil {
		return server.CreateTunnel500JSONResponse(*s.captureDetail(f)), nil
	}

	return server.CreateTunnel201JSONResponse(*convert.Tunnel(t)), nil
}
