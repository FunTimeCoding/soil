package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelTerminations(
	_ context.Context,
	_ server.ListTunnelTerminationsRequestObject,
) (server.ListTunnelTerminationsResponseObject, error) {
	terminations, e := s.service.TunnelTerminations()

	if forbidden.Is(e) {
		return server.ListTunnelTerminations403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListTunnelTerminations500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListTunnelTerminations200JSONResponse(
			convert.TunnelTerminations(terminations),
		),
		nil
}
