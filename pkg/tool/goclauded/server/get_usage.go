package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetUsage(
	_ context.Context,
	_ server.GetUsageRequestObject,
) (server.GetUsageResponseObject, error) {
	result := s.service.Usage()

	if result == nil {
		return server.GetUsage204Response{}, nil
	}

	return server.GetUsage200JSONResponse{
		FiveHourPercent: result.FiveHourPercent,
		FiveHourReset:   result.FiveHourReset,
		SevenDayPercent: result.SevenDayPercent,
		SevenDayReset:   result.SevenDayReset,
		FablePercent:    result.FablePercent,
		FableReset:      result.FableReset,
		LastUpdated:     result.LastUpdated,
	}, nil
}
