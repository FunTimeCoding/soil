package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/convert"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/server"
)

func (s *Server) PostGroup(
	_ context.Context,
	r server.PostGroupRequestObject,
) (server.PostGroupResponseObject, error) {
	result, e := s.service.CreateGroup(r.Body.Name)

	if validation.Is(e) {
		return server.PostGroup400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.PostGroup403JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostGroup409JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.PostGroup503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PostGroup500JSONResponse(*s.answer(e)), nil
	}

	return server.PostGroup200JSONResponse(*convert.Group(result)), nil
}
