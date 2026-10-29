package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/server"
)

func (s *Server) DeleteHost(
	_ context.Context,
	r server.DeleteHostRequestObject,
) (server.DeleteHostResponseObject, error) {
	e := s.service.DeleteHost(r.Identifier, apply(r.Params.Apply))

	if not_found.Is(e) {
		return server.DeleteHost404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteHost500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteHost204Response{}, nil
}
