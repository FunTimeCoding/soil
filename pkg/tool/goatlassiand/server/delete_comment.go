package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeleteComment(
	x context.Context,
	r server.DeleteCommentRequestObject,
) (server.DeleteCommentResponseObject, error) {
	if e := s.service.DeleteComment(r.Key, r.Identifier); e != nil {
		return server.DeleteComment500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteComment204Response{}, nil
}
