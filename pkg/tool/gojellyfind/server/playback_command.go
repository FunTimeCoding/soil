package server

import (
	"context"
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

	if e := s.client.PlaybackCommand(
		r.Identifier,
		r.Body.Command,
		seek,
	); e != nil {
		return server.PlaybackCommand500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.PlaybackCommand204Response{}, nil
}
