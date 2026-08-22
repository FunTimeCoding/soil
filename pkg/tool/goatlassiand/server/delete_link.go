package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeleteLink(
	x context.Context,
	r server.DeleteLinkRequestObject,
) (server.DeleteLinkResponseObject, error) {
	if e := s.service.DeleteLink(r.Identifier); e != nil {
		return server.DeleteLink500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.DeleteLink204Response{}, nil
}
