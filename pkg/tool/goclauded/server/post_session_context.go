package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
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

	s.service.RecordContext(
		r.Identifier,
		r.Body.UsedPercentage,
		window,
		model,
	)

	if r.Body.FiveHourPercent == nil || r.Body.SevenDayPercent == nil {
		return server.PostSessionContext200Response{}, nil
	}

	if e := s.service.RecordRateLimits(
		*r.Body.FiveHourPercent,
		*r.Body.SevenDayPercent,
		epochTime(r.Body.FiveHourReset),
		epochTime(r.Body.SevenDayReset),
	); e != nil {
		return server.PostSessionContext500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostSessionContext200Response{}, nil
}
