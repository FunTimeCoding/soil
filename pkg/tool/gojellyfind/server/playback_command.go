package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) PlaybackCommand(
	_ context.Context,
	r server.PlaybackCommandRequestObject,
) (server.PlaybackCommandResponseObject, error) {
	var seek int64

	if r.Body.SeekPositionTicks != nil {
		seek = *r.Body.SeekPositionTicks
	}

	e := s.service.PlaybackCommand(r.Identifier, r.Body.Command, seek)

	if validation.Is(e) {
		return server.PlaybackCommand400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.PlaybackCommand404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.PlaybackCommand500JSONResponse(*s.answer(e)), nil
	}

	return server.PlaybackCommand204Response{}, nil
}
