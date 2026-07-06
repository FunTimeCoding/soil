package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListSpaces(
	_ context.Context,
	_ server.ListSpacesRequestObject,
) (server.ListSpacesResponseObject, error) {
	result, e := s.confluence.Spaces()

	if e != nil {
		return server.ListSpaces500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListSpaces200JSONResponse(
		convert.ConfluenceSpaces(result),
	), nil
}
