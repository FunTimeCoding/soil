package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) ScoreTask(
	_ context.Context,
	r server.ScoreTaskRequestObject,
) (server.ScoreTaskResponseObject, error) {
	result, e := s.habitica.Score(
		r.Identifier,
		string(r.Direction),
	)

	if e != nil {
		return server.ScoreTask500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.ScoreTask200JSONResponse(
		*convert.ScoreResult(result),
	), nil
}
