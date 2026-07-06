package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetStats(
	_ context.Context,
	_ server.GetStatsRequestObject,
) (server.GetStatsResponseObject, error) {
	result, e := s.habitica.UserStats()

	if e != nil {
		return server.GetStats500JSONResponse(
			*s.captureDetail(e),
		), nil
	}

	return server.GetStats200JSONResponse(*convert.Stats(result)), nil
}
