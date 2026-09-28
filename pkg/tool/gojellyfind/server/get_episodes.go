package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetEpisodes(
	_ context.Context,
	r server.GetEpisodesRequestObject,
) (server.GetEpisodesResponseObject, error) {
	items, e := s.client.Episodes(r.Identifier)

	if e != nil {
		return server.GetEpisodes500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetEpisodes200JSONResponse(convertItems(items)), nil
}
