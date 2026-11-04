package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetUsage(
	x context.Context,
	_ server.GetUsageRequestObject,
) (server.GetUsageResponseObject, error) {
	result, found, e := s.usage.FindUsage()

	if e != nil {
		return server.GetUsage500JSONResponse(
			*s.captureFail(x, e, soil.UnexpectedError),
		), nil
	}

	if !found {
		return server.GetUsage204Response{}, nil
	}

	response := server.GetUsage200JSONResponse{
		FiveHourPercent: result.FiveHourPercent,
		FiveHourReset:   result.FiveHourReset,
		SevenDayPercent: result.SevenDayPercent,
		SevenDayReset:   result.SevenDayReset,
		FablePercent:    result.FablePercent,
		FableReset:      result.FableReset,
		LastUpdated:     result.LastUpdated,
	}

	if !result.FableResetAt.IsZero() {
		response.FableResetAt = &result.FableResetAt
	}

	return response, nil
}
