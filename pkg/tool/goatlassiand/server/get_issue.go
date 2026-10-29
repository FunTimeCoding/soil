package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) GetIssue(
	_ context.Context,
	r server.GetIssueRequestObject,
) (server.GetIssueResponseObject, error) {
	result, e := s.service.Issue(r.Key)

	if forbidden.Is(e) {
		return server.GetIssue403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetIssue404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetIssue503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetIssue500JSONResponse(*s.answer(e)), nil
	}

	converted := convert.JiraIssue(result)

	if r.Params.Comments != nil && *r.Params.Comments {
		converted.Comments = convert.JiraComments(result)
	}

	return server.GetIssue200JSONResponse(*converted), nil
}
