package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) LinkIssues(
	x context.Context,
	r server.LinkIssuesRequestObject,
) (server.LinkIssuesResponseObject, error) {
	linkType := "Relates"

	if r.Body.LinkType != nil {
		linkType = *r.Body.LinkType
	}

	if e := s.service.LinkIssues(r.Key, r.Body.TargetKey, linkType); e != nil {
		return server.LinkIssues500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.LinkIssues204Response{}, nil
}
