package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) AddPageComment(
	_ context.Context,
	r server.AddPageCommentRequestObject,
) (server.AddPageCommentResponseObject, error) {
	e := s.service.AddPageComment(r.Identifier, r.Body.Body)

	if forbidden.Is(e) {
		return server.AddPageComment403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.AddPageComment404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.AddPageComment503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AddPageComment500JSONResponse(*s.answer(e)), nil
	}

	return server.AddPageComment204Response{}, nil
}
