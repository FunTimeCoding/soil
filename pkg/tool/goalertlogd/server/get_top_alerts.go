package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"time"
)

func (s *Server) GetTopAlerts(
	_ context.Context,
	r server.GetTopAlertsRequestObject,
) (server.GetTopAlertsResponseObject, error) {
	n := 25
	now := time.Now()
	start := now.Add(-7 * 24 * time.Hour)
	end := now

	if r.Params.N != nil {
		n = *r.Params.N
	}

	if r.Params.Start != nil {
		start = *r.Params.Start
	}

	if r.Params.End != nil {
		end = *r.Params.End
	}

	records, e := s.service.TopAlerts(n, start, end)

	if e != nil {
		return server.GetTopAlerts500JSONResponse(*s.answer(e)), nil
	}

	return server.GetTopAlerts200JSONResponse(convert.TopAlerts(records)), nil
}
