package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) DeleteMemory(
	_ context.Context,
	r server.DeleteMemoryRequestObject,
) (server.DeleteMemoryResponseObject, error) {
	source := ""

	if r.Params.Source != nil {
		source = *r.Params.Source
	}

	e := s.service.ForgetMemory(r.Identifier, source)

	if not_found.Is(e) {
		return server.DeleteMemory404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteMemory500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteMemory200JSONResponse{Identifier: r.Identifier}, nil
}
