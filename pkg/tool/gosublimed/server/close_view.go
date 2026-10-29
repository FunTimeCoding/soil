package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/server"
)

func (s *Server) CloseView(
	_ context.Context,
	r server.CloseViewRequestObject,
) (server.CloseViewResponseObject, error) {
	e := s.service.CloseView(r.Id)

	if not_found.Is(e) {
		return server.CloseView404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.CloseView500JSONResponse(*s.answer(e)), nil
	}

	return server.CloseView204Response{}, nil
}
