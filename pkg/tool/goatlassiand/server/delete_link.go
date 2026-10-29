package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeleteLink(
	_ context.Context,
	r server.DeleteLinkRequestObject,
) (server.DeleteLinkResponseObject, error) {
	e := s.service.DeleteLink(r.Identifier)

	if forbidden.Is(e) {
		return server.DeleteLink403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.DeleteLink404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeleteLink503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteLink500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteLink204Response{}, nil
}
