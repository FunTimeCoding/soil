package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) ListProjects(
	_ context.Context,
	_ server.ListProjectsRequestObject,
) (server.ListProjectsResponseObject, error) {
	result, e := s.service.Projects()

	if forbidden.Is(e) {
		return server.ListProjects403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ListProjects503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListProjects500JSONResponse(*s.answer(e)), nil
	}

	return server.ListProjects200JSONResponse(convert.JiraProjects(result)), nil
}
