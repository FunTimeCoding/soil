package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetStats(
	_ context.Context,
	_ server.GetStatsRequestObject,
) (server.GetStatsResponseObject, error) {
	result, e := s.service.UserStats()

	if forbidden.Is(e) {
		return server.GetStats403JSONResponse{Error: e.Error()}, nil
	}

	if unavailable.Is(e) {
		return server.GetStats503JSONResponse{Error: e.Error()}, nil
	}

	if e != nil {
		return server.GetStats500JSONResponse(*s.answer(e)), nil
	}

	return server.GetStats200JSONResponse(*convert.Stats(result)), nil
}
