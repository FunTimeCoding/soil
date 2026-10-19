package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
)

func (s *Server) PostMemory(
	_ context.Context,
	r server.PostMemoryRequestObject,
) (server.PostMemoryResponseObject, error) {
	m, e := s.service.CreateMemory(saveOption(*r.Body))

	if e != nil {
		if validation.Is(e) {
			return server.PostMemory400JSONResponse{Error: e.Error()}, nil
		}

		return server.PostMemory500JSONResponse(*s.answer(e)), nil
	}

	return server.PostMemory200JSONResponse{Identifier: m.Identifier}, nil
}
