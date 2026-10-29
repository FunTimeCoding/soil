package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/server"
)

func (s *Server) ReloadProcfile(
	_ context.Context,
	_ server.ReloadProcfileRequestObject,
) (server.ReloadProcfileResponseObject, error) {
	e := s.service.ReloadProcfile()

	if validation.Is(e) {
		return server.ReloadProcfile400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ReloadProcfile500JSONResponse(*s.answer(e)), nil
	}

	return server.ReloadProcfile204Response{}, nil
}
