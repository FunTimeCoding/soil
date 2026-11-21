package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
)

func (s *Server) GetScoreStatus(
	context.Context,
	server.GetScoreStatusRequestObject,
) (server.GetScoreStatusResponseObject, error) {
	return server.GetScoreStatus200JSONResponse(
		*convert.ScoreStatus(s.score.Status()),
	), nil
}
