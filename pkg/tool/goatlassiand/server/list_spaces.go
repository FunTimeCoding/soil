package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListSpaces(
	_ context.Context,
	_ server.ListSpacesRequestObject,
) (server.ListSpacesResponseObject, error) {
	result, e := s.service.Spaces()

	if forbidden.Is(e) {
		return server.ListSpaces403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ListSpaces503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSpaces500JSONResponse(*s.answer(e)), nil
	}

	return server.ListSpaces200JSONResponse(convert.ConfluenceSpaces(result)), nil
}
