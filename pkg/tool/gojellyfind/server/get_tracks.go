package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetTracks(
	_ context.Context,
	r server.GetTracksRequestObject,
) (server.GetTracksResponseObject, error) {
	items, e := s.service.Tracks(r.Identifier)

	if validation.Is(e) {
		return server.GetTracks400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetTracks500JSONResponse(*s.answer(e)), nil
	}

	return server.GetTracks200JSONResponse(convertItems(items)), nil
}
