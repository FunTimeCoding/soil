package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) UpdateComment(
	x context.Context,
	r server.UpdateCommentRequestObject,
) (server.UpdateCommentResponseObject, error) {
	if e := s.service.UpdateComment(
		r.Key,
		r.Identifier,
		r.Body.Body,
	); e != nil {
		return server.UpdateComment500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.UpdateComment204Response{}, nil
}
