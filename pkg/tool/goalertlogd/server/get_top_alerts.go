package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
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

	records, e := s.store.Top(n, start, end)

	if e != nil {
		return server.GetTopAlerts500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	result := make(server.GetTopAlerts200JSONResponse, 0, len(records))

	for _, c := range records {
		result = append(
			result,
			server.TopAlertsResponse{
				Name:            c.Name,
				Count:           c.Count,
				AverageDuration: c.AverageDuration.String(),
				CurrentlyFiring: c.CurrentlyFiring,
				Severity:        c.Severity,
			},
		)
	}

	return result, nil
}
