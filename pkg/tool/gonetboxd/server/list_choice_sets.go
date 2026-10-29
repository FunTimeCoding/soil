package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) ListChoiceSets(
	_ context.Context,
	_ server.ListChoiceSetsRequestObject,
) (server.ListChoiceSetsResponseObject, error) {
	result, e := s.service.ChoiceSets()

	if forbidden.Is(e) {
		return server.ListChoiceSets403JSONResponse(*clientError(e)), nil
	}

	if e != nil {
		return server.ListChoiceSets500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.ListChoiceSets200JSONResponse(convert.ChoiceSets(result)), nil
}
