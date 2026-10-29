package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) RestartProcess(
	_ context.Context,
	r server.RestartProcessRequestObject,
) (server.RestartProcessResponseObject, error) {
	e := s.service.Restart(r.Name)

	if not_found.Is(e) {
		return server.RestartProcess404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.RestartProcess500JSONResponse(*s.answer(e)), nil
	}

	return server.RestartProcess204Response{}, nil
}
