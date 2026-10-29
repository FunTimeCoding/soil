package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) SearchIssues(
	_ context.Context,
	r server.SearchIssuesRequestObject,
) (server.SearchIssuesResponseObject, error) {
	result, e := s.service.SearchIssues(r.Params.Query, r.Params.Limit)

	if validation.Is(e) {
		return server.SearchIssues400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.SearchIssues403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.SearchIssues503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SearchIssues500JSONResponse(*s.answer(e)), nil
	}

	return server.SearchIssues200JSONResponse(convert.JiraIssues(result)), nil
}
