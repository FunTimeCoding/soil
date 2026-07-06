package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnels(
	_ context.Context,
	_ server.ListTunnelsRequestObject,
) (server.ListTunnelsResponseObject, error) {
	tunnels, e := s.client.Tunnels()

	if e != nil {
		return server.ListTunnels500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListTunnels200JSONResponse(
		convert.Tunnels(tunnels),
	), nil
}
