package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) ScoreTask(
	_ context.Context,
	r server.ScoreTaskRequestObject,
) (server.ScoreTaskResponseObject, error) {
	result, e := s.service.Score(r.Identifier, string(r.Direction))

	if not_found.Is(e) {
		return server.ScoreTask404JSONResponse{Error: e.Error()}, nil
	}

	if forbidden.Is(e) {
		return server.ScoreTask403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.ScoreTask503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ScoreTask500JSONResponse(*s.answer(e)), nil
	}

	return server.ScoreTask200JSONResponse(*convert.ScoreResult(result)), nil
}
