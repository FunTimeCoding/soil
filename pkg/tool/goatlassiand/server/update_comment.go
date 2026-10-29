package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) UpdateComment(
	_ context.Context,
	r server.UpdateCommentRequestObject,
) (server.UpdateCommentResponseObject, error) {
	e := s.service.UpdateComment(r.Key, r.Identifier, r.Body.Body)

	if validation.Is(e) {
		return server.UpdateComment400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.UpdateComment403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.UpdateComment404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.UpdateComment503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.UpdateComment500JSONResponse(*s.answer(e)), nil
	}

	return server.UpdateComment204Response{}, nil
}
