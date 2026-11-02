package server

import (
	"context"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostSessionContext(
	_ context.Context,
	r server.PostSessionContextRequestObject,
) (server.PostSessionContextResponseObject, error) {
	window := 0

	if r.Body.WindowSize != nil {
		window = *r.Body.WindowSize
	}

	model := ""

	if r.Body.Model != nil {
		model = *r.Body.Model
	}

	s.roster.RecordContext(r.Identifier, r.Body.UsedPercentage, window, model)

	if r.Body.FablePercent != nil {
		if e := s.usage.RecordFable(
			*r.Body.FablePercent,
			"",
			epochTimePointer(r.Body.FableReset),
		); e != nil {
			return server.PostSessionContext500JSONResponse(
				*s.captureFail(e, soil.UnexpectedError),
			), nil
		}
	}

	if r.Body.SevenDayPercent == nil {
		return server.PostSessionContext200Response{}, nil
	}

	fiveHour := 0

	if r.Body.FiveHourPercent != nil {
		fiveHour = *r.Body.FiveHourPercent
	}

	if e := s.usage.RecordRateLimits(
		fiveHour,
		*r.Body.SevenDayPercent,
		epochTime(r.Body.FiveHourReset),
		epochTime(r.Body.SevenDayReset),
	); e != nil {
		return server.PostSessionContext500JSONResponse(
			*s.captureFail(e, soil.UnexpectedError),
		), nil
	}

	return server.PostSessionContext200Response{}, nil
}
