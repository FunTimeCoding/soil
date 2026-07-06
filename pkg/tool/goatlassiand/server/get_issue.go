package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetIssue(
	_ context.Context,
	r server.GetIssueRequestObject,
) (server.GetIssueResponseObject, error) {
	result, e := s.jira.Issue(r.Key)

	if e != nil {
		return server.GetIssue500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetIssue200JSONResponse(
		*convert.JiraIssue(result),
	), nil
}
