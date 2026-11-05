package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelGroups(
	x context.Context,
	_ server.ListTunnelGroupsRequestObject,
) (server.ListTunnelGroupsResponseObject, error) {
	groups, e := s.service.TunnelGroups()

	if unavailable.Is(e) {
		return server.ListTunnelGroups503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListTunnelGroups500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListTunnelGroups200JSONResponse(convert.TunnelGroups(groups)), nil
}
