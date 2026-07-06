package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
)

func toResponse(records []store.Record) []server.AlertsResponse {
	result := make([]server.AlertsResponse, 0, len(records))

	for _, record := range records {
		entry := server.AlertsResponse{
			Fingerprint: record.Fingerprint,
			Name:        record.Name,
			Severity:    record.Severity,
			Summary:     record.Summary,
			Labels:      record.Labels,
			Start:       record.Start.Format(DateFormat),
			Status:      server.Firing,
		}

		if record.End != nil {
			entry.End = new(record.End.Format(DateFormat))
			entry.Status = server.Resolved
		}

		result = append(result, entry)
	}

	return result
}
