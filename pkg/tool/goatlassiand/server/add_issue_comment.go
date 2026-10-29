package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) AddIssueComment(
	_ context.Context,
	r server.AddIssueCommentRequestObject,
) (server.AddIssueCommentResponseObject, error) {
	e := s.service.Comment(r.Key, r.Body.Body)

	if validation.Is(e) {
		return server.AddIssueComment400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.AddIssueComment403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.AddIssueComment404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.AddIssueComment503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.AddIssueComment500JSONResponse(*s.answer(e)), nil
	}

	return server.AddIssueComment204Response{}, nil
}
