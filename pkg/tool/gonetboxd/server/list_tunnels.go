package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnels(
	x context.Context,
	_ server.ListTunnelsRequestObject,
) (server.ListTunnelsResponseObject, error) {
	tunnels, e := s.service.Tunnels()

	if unavailable.Is(e) {
		return server.ListTunnels503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListTunnels500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListTunnels200JSONResponse(convert.Tunnels(tunnels)), nil
}
