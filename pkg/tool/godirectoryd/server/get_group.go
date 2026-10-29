package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) GetGroup(
	_ context.Context,
	_ server.GetGroupRequestObject,
) (server.GetGroupResponseObject, error) {
	result, e := s.service.Groups()

	if forbidden.Is(e) {
		return server.GetGroup403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetGroup503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetGroup500JSONResponse(*s.answer(e)), nil
	}

	return server.GetGroup200JSONResponse(convert.Groups(result)), nil
}
