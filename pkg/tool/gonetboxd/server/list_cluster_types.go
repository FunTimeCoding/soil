package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListClusterTypes(
	x context.Context,
	_ server.ListClusterTypesRequestObject,
) (server.ListClusterTypesResponseObject, error) {
	types, e := s.service.ClusterTypes()

	if unavailable.Is(e) {
		return server.ListClusterTypes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListClusterTypes500JSONResponse(
			*s.captureDetail(x, e),
		), nil
	}

	return server.ListClusterTypes200JSONResponse(convert.ClusterTypes(types)), nil
}
