package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func (s *Server) GetTracks(
	_ context.Context,
	r server.GetTracksRequestObject,
) (server.GetTracksResponseObject, error) {
	items, e := s.client.Tracks(r.Identifier)

	if e != nil {
		return server.GetTracks500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.GetTracks200JSONResponse(convertItems(items)), nil
}
