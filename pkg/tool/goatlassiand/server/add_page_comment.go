package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) AddPageComment(
	_ context.Context,
	r server.AddPageCommentRequestObject,
) (server.AddPageCommentResponseObject, error) {
	if e := s.confluence.AddComment(
		r.Identifier,
		r.Body.Body,
	); e != nil {
		return server.AddPageComment500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.AddPageComment204Response{}, nil
}
