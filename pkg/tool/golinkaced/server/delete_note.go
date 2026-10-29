package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) DeleteNote(
	_ context.Context,
	r server.DeleteNoteRequestObject,
) (server.DeleteNoteResponseObject, error) {
	e := s.service.DeleteNote(int(r.Identifier))

	if not_found.Is(e) {
		return server.DeleteNote404JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.DeleteNote403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.DeleteNote503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.DeleteNote500JSONResponse(*s.answer(e)), nil
	}

	return server.DeleteNote200Response{}, nil
}
