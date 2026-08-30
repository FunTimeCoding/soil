package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
)

func TopAlerts(records []store.TopRecord) []server.TopAlertsResponse {
	result := make([]server.TopAlertsResponse, 0, len(records))

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

	return result
}
