package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) SetVolume(
	_ context.Context,
	r server.SetVolumeRequestObject,
) (server.SetVolumeResponseObject, error) {
	e := s.client.SetVolume(r.Identifier, r.Body.Level)

	if not_found.Is(e) {
		return server.SetVolume404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.SetVolume500JSONResponse(*s.answer(e)), nil
	}

	return server.SetVolume204Response{}, nil
}
