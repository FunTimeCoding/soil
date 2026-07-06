package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelTerminations(
	_ context.Context,
	_ server.ListTunnelTerminationsRequestObject,
) (server.ListTunnelTerminationsResponseObject, error) {
	terminations, e := s.client.TunnelTerminations()

	if e != nil {
		return server.ListTunnelTerminations500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListTunnelTerminations200JSONResponse(
		convert.TunnelTerminations(terminations),
	), nil
}
