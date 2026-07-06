package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"time"
)

func (s *Server) GetRecentAlerts(
	_ context.Context,
	r server.GetRecentAlertsRequestObject,
) (server.GetRecentAlertsResponseObject, error) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if r.Params.Start != nil {
		start = *r.Params.Start
	}

	if r.Params.End != nil {
		end = *r.Params.End
	}

	records, e := s.store.ByTimeRange(start, end)

	if e != nil {
		return server.GetRecentAlerts500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.GetRecentAlerts200JSONResponse(toResponse(records)), nil
}
