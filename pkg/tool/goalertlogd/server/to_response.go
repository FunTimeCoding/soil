package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
)

func toResponse(records []record.Record) []server.AlertsResponse {
	result := make([]server.AlertsResponse, 0, len(records))

	for _, r := range records {
		entry := server.AlertsResponse{
			Fingerprint: r.Fingerprint,
			Name:        r.Name,
			Severity:    r.Severity,
			Summary:     r.Summary,
			Labels:      r.Labels,
			Start:       r.Start.Format(DateFormat),
			Status:      server.Firing,
		}

		if r.End != nil {
			entry.End = new(r.End.Format(DateFormat))
			entry.Status = server.Resolved
		}

		result = append(result, entry)
	}

	return result
}
