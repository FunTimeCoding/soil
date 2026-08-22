package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeletePage(
	_ context.Context,
	r server.DeletePageRequestObject,
) (server.DeletePageResponseObject, error) {
	draft := r.Params.Draft != nil && *r.Params.Draft

	if draft {
		if e := s.confluence.DeleteDraft(r.Identifier); e != nil {
			return server.DeletePage500JSONResponse(*s.captureDetail(e)), nil
		}
	} else {
		if e := s.confluence.Delete(r.Identifier); e != nil {
			return server.DeletePage500JSONResponse(*s.captureDetail(e)), nil
		}
	}

	return server.DeletePage204Response{}, nil
}
