package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func (s *Server) GetDiff(
	_ context.Context,
	r server.GetDiffRequestObject,
) (server.GetDiffResponseObject, error) {
	result, e := s.service.DiffSymbols(
		r.Params.Module,
		text(r.Params.From),
		text(r.Params.To),
	)

	if not_found.Is(e) {
		return server.GetDiff404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetDiff500JSONResponse(*s.answer(e)), nil
	}

	return server.GetDiff200JSONResponse(diffResponse(result)), nil
}
