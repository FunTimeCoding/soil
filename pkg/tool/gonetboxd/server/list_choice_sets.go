package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListChoiceSets(
	_ context.Context,
	_ server.ListChoiceSetsRequestObject,
) (server.ListChoiceSetsResponseObject, error) {
	result, e := s.client.ChoiceSets()

	if e != nil {
		return server.ListChoiceSets500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListChoiceSets200JSONResponse(convert.ChoiceSets(result)), nil
}
