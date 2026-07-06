package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) TransitionIssue(
	_ context.Context,
	r server.TransitionIssueRequestObject,
) (server.TransitionIssueResponseObject, error) {
	if e := s.jira.Transition(
		r.Key,
		r.Body.TransitionIdentifier,
	); e != nil {
		return server.TransitionIssue500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.TransitionIssue204Response{}, nil
}
