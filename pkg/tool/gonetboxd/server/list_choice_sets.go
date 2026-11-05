package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListChoiceSets(
	x context.Context,
	_ server.ListChoiceSetsRequestObject,
) (server.ListChoiceSetsResponseObject, error) {
	result, e := s.service.ChoiceSets()

	if unavailable.Is(e) {
		return server.ListChoiceSets503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListChoiceSets500JSONResponse(*s.captureDetail(x, e)), nil
	}

	return server.ListChoiceSets200JSONResponse(convert.ChoiceSets(result)), nil
}
