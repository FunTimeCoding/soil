package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) GetUser(
	_ context.Context,
	_ server.GetUserRequestObject,
) (server.GetUserResponseObject, error) {
	result, e := s.service.Users()

	if forbidden.Is(e) {
		return server.GetUser403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetUser503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetUser500JSONResponse(*s.answer(e)), nil
	}

	return server.GetUser200JSONResponse(convert.Users(result)), nil
}
