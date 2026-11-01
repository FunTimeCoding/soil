package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/server"
)

func (s *Server) PostRun(
	_ context.Context,
	r server.PostRunRequestObject,
) (server.PostRunResponseObject, error) {
	e := s.service.Run(r.Body.Intent)

	if e == nil {
		return server.PostRun202Response{}, nil
	}

	if validation.Is(e) {
		return server.PostRun400JSONResponse{Error: e.Error()}, nil
	}

	if conflict.Is(e) {
		return server.PostRun409JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.PostRun503JSONResponse{Error: e.Error()}, nil
	}

	return server.PostRun500JSONResponse(*s.answer(e)), nil
}
