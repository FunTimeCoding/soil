package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
)

func (s *Server) DeleteEntry(
	_ context.Context,
	r server.DeleteEntryRequestObject,
) (server.DeleteEntryResponseObject, error) {
	e := s.service.Delete(uint(r.Id))

	if not_found.Is(e) {
		return server.DeleteEntry404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteEntry500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteEntry204Response{}, nil
}
