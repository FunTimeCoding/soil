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
		SessionPercent:      result.SessionPercent,
		SessionReset:        result.SessionReset,
		WeeklyAllPercent:    result.WeeklyAllPercent,
		WeeklyAllReset:      result.WeeklyAllReset,
		WeeklySonnetPercent: result.WeeklySonnetPercent,
		WeeklyDesignPercent: result.WeeklyDesignPercent,
		RoutineRuns:         result.RoutineRuns,
		CreditSpent:         result.CreditSpent,
		CreditReset:         result.CreditReset,
		CreditPercent:       result.CreditPercent,
		LastUpdated:         result.LastUpdated,
	}, nil
}
