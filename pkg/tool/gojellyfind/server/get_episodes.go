package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetEpisodes(
	x context.Context,
	r server.GetEpisodesRequestObject,
) (server.GetEpisodesResponseObject, error) {
	items, e := s.service.Episodes(r.Identifier)

	if validation.Is(e) {
		return server.GetEpisodes400JSONResponse{Error: e.Error()}, nil
	}

	if not_found.Is(e) {
		return server.GetEpisodes404JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetEpisodes503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetEpisodes500JSONResponse(*s.answer(x, e)), nil
	}

	return server.GetEpisodes200JSONResponse(convertItems(items)), nil
}
