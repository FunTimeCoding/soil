package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) ListSystems(
	x context.Context,
	r server.ListSystemsRequestObject,
) (server.ListSystemsResponseObject, error) {
	result, e := s.system.Summaries(convert.Minimum(r.Params.Score))

	if validation.Is(e) {
		return server.ListSystems400JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ListSystems500JSONResponse(*s.answer(x, e)), nil
	}

	return server.ListSystems200JSONResponse(
		convert.SystemSummaries(result),
	), nil
}
