package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) StopProcess(
	_ context.Context,
	r server.StopProcessRequestObject,
) (server.StopProcessResponseObject, error) {
	e := s.service.Stop(r.Name)

	if not_found.Is(e) {
		return server.StopProcess404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.StopProcess500JSONResponse(*s.answer(e)), nil
	}

	return server.StopProcess204Response{}, nil
}
