package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func (s *Server) DeletePage(
	_ context.Context,
	r server.DeletePageRequestObject,
) (server.DeletePageResponseObject, error) {
	draft := r.Params.Draft != nil && *r.Params.Draft
	e := s.service.DeletePage(r.Identifier, draft)

	if forbidden.Is(e) {
		return server.DeletePage403JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.DeletePage404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeletePage503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeletePage500JSONResponse(*s.answer(e)), nil
	}

	return server.DeletePage204Response{}, nil
}
