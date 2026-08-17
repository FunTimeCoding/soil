package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListInterfaces(
	_ context.Context,
	_ server.ListInterfacesRequestObject,
) (server.ListInterfacesResponseObject, error) {
	result, e := s.opnsense.Interfaces()

	if e != nil {
		return server.ListInterfaces500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListInterfaces200JSONResponse(
		convert.NetworkInterfaces(result),
	), nil
}
