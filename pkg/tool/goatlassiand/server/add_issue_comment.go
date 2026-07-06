package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) AddIssueComment(
	_ context.Context,
	r server.AddIssueCommentRequestObject,
) (server.AddIssueCommentResponseObject, error) {
	if e := s.jira.Comment(r.Key, r.Body.Body); e != nil {
		return server.AddIssueComment500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.AddIssueComment204Response{}, nil
}
