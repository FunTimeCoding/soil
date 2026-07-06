package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchIssues(
	_ context.Context,
	r server.SearchIssuesRequestObject,
) (server.SearchIssuesResponseObject, error) {
	if r.Params.Limit != nil {
		result, e := s.jira.SearchLimit(*r.Params.Limit, r.Params.Query)

		if e != nil {
			return server.SearchIssues500JSONResponse(
				*s.captureDetail(e),
			), nil
		}

		return server.SearchIssues200JSONResponse(
			convert.JiraIssues(result),
		), nil
	}

	result, e := s.jira.Search(r.Params.Query)

	if e != nil {
		return server.SearchIssues500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.SearchIssues200JSONResponse(
		convert.JiraIssues(result),
	), nil
}
