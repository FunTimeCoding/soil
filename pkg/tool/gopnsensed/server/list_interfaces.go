package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/convert"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) ListInterfaces(
	x context.Context,
	_ server.ListInterfacesRequestObject,
) (server.ListInterfacesResponseObject, error) {
	result, e := s.service.Interfaces()

	if e != nil {
		return server.ListInterfaces500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListInterfaces200JSONResponse(
		convert.NetworkInterfaces(result),
	), nil
}
