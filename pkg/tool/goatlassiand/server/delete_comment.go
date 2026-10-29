package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeleteComment(
	_ context.Context,
	r server.DeleteCommentRequestObject,
) (server.DeleteCommentResponseObject, error) {
	e := s.service.DeleteComment(r.Key, r.Identifier)

	if forbidden.Is(e) {
		return server.DeleteComment403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.DeleteComment404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeleteComment503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteComment500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteComment204Response{}, nil
}
