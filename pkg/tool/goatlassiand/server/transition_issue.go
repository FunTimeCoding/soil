package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) TransitionIssue(
	_ context.Context,
	r server.TransitionIssueRequestObject,
) (server.TransitionIssueResponseObject, error) {
	e := s.service.Transition(r.Key, r.Body.TransitionIdentifier)

	if validation.Is(e) {
		return server.TransitionIssue400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.TransitionIssue403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.TransitionIssue404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.TransitionIssue503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.TransitionIssue500JSONResponse(*s.answer(e)), nil
	}

	return server.TransitionIssue204Response{}, nil
}
