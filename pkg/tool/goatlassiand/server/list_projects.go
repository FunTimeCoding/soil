package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListProjects(
	_ context.Context,
	_ server.ListProjectsRequestObject,
) (server.ListProjectsResponseObject, error) {
	result, e := s.jira.Projects()

	if e != nil {
		return server.ListProjects500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ListProjects200JSONResponse(
		convert.JiraProjects(result),
	), nil
}
