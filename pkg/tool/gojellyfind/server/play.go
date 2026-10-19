package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) Play(
	_ context.Context,
	r server.PlayRequestObject,
) (server.PlayResponseObject, error) {
	command := "PlayNow"

	if r.Body.PlayCommand != nil {
		command = *r.Body.PlayCommand
	}

	var start int64

	if r.Body.StartPositionTicks != nil {
		start = *r.Body.StartPositionTicks
	}

	e := s.client.Play(r.Identifier, r.Body.ItemIds, command, start)

	if not_found.Is(e) {
		return server.Play404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.Play500JSONResponse(*s.answer(e)), nil
	}

	return server.Play204Response{}, nil
}
