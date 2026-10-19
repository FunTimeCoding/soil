package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) ScoreTask(
	_ context.Context,
	r server.ScoreTaskRequestObject,
) (server.ScoreTaskResponseObject, error) {
	result, e := s.habitica.Score(r.Identifier, string(r.Direction))

	if not_found.Is(e) {
		return server.ScoreTask404JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.ScoreTask500JSONResponse(*s.answer(e)), nil
	}

	return server.ScoreTask200JSONResponse(*convert.ScoreResult(result)), nil
}
