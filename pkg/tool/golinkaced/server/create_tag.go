package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) CreateTag(
	_ context.Context,
	r server.CreateTagRequestObject,
) (server.CreateTagResponseObject, error) {
	result, f := s.service.CreateTag(r.Body.Name)

	if validation.Is(f) {
		return server.CreateTag400JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.CreateTag403JSONResponse{Error: f.Error()}, nil
	}

	if conflict.Is(f) {
		return server.CreateTag409JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.CreateTag503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.CreateTag500JSONResponse(*s.answer(f)), nil
	}

	return server.CreateTag200JSONResponse{
		Identifier: int32(result.Identifier),
		Name:       result.Name,
	}, nil
}
