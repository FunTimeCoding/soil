package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) LinkIssues(
	_ context.Context,
	r server.LinkIssuesRequestObject,
) (server.LinkIssuesResponseObject, error) {
	linkType := "Relates"

	if r.Body.LinkType != nil {
		linkType = *r.Body.LinkType
	}

	e := s.service.LinkIssues(r.Key, r.Body.TargetKey, linkType)

	if validation.Is(e) {
		return server.LinkIssues400JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.LinkIssues403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.LinkIssues404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.LinkIssues503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.LinkIssues500JSONResponse(*s.answer(e)), nil
	}

	return server.LinkIssues204Response{}, nil
}
