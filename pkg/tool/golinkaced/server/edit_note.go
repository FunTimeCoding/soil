package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/server"
)

func (s *Server) EditNote(
	_ context.Context,
	r server.EditNoteRequestObject,
) (server.EditNoteResponseObject, error) {
	result, f := s.service.UpdateNote(int(r.Identifier), r.Body.Text)

	if not_found.Is(f) {
		return server.EditNote404JSONResponse{Error: f.Error()}, nil
	}

	if validation.Is(f) {
		return server.EditNote400JSONResponse{Error: f.Error()}, nil
	}

	if forbidden.Is(f) {
		return server.EditNote403JSONResponse{Error: f.Error()}, nil
	}

	if unavailable.Is(f) {
		return server.EditNote503JSONResponse{Error: f.Error()}, nil
	}

	if f != nil {
		return server.EditNote500JSONResponse(*s.answer(f)), nil
	}

	return server.EditNote200JSONResponse{
		Identifier: int32(result.Identifier),
		Text:       result.Text,
	}, nil
}
