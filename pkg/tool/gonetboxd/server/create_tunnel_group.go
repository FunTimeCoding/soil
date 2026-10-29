package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreateTunnelGroup(
	_ context.Context,
	r server.CreateTunnelGroupRequestObject,
) (server.CreateTunnelGroupResponseObject, error) {
	g, e := s.service.CreateTunnelGroup(r.Body.Name)

	if validation.Is(e) {
		return server.CreateTunnelGroup400JSONResponse(*clientError(e)), nil
	}

	if forbidden.Is(e) {
		return server.CreateTunnelGroup403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.CreateTunnelGroup500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreateTunnelGroup201JSONResponse(*convert.TunnelGroup(g)), nil
}
