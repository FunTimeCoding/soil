package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListSystems(
	x context.Context,
	_ server.ListSystemsRequestObject,
) (server.ListSystemsResponseObject, error) {
	result, e := s.system.Summaries()

	if e != nil {
		return server.ListSystems500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListSystems200JSONResponse(
		convert.SystemSummaries(result),
	), nil
}
