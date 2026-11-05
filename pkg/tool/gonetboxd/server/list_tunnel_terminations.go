package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelTerminations(
	x context.Context,
	_ server.ListTunnelTerminationsRequestObject,
) (server.ListTunnelTerminationsResponseObject, error) {
	terminations, e := s.service.TunnelTerminations()

	if unavailable.Is(e) {
		return server.ListTunnelTerminations503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListTunnelTerminations500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListTunnelTerminations200JSONResponse(
			convert.TunnelTerminations(terminations),
		),
		nil
}
