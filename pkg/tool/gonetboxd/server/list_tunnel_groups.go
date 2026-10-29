package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListTunnelGroups(
	_ context.Context,
	_ server.ListTunnelGroupsRequestObject,
) (server.ListTunnelGroupsResponseObject, error) {
	groups, e := s.service.TunnelGroups()

	if forbidden.Is(e) {
		return server.ListTunnelGroups403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListTunnelGroups500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListTunnelGroups200JSONResponse(convert.TunnelGroups(groups)), nil
}
