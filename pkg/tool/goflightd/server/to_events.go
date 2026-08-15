package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
)

func toEvents(v []event.Event) []server.EventResponse {
	result := make([]server.EventResponse, 0, len(v))

	for _, w := range v {
		result = append(
			result,
			server.EventResponse{
				Time:      w.Time.Format(constant.DateFormat),
				Process:   w.Process,
				Subsystem: w.Subsystem,
				Category:  w.Category,
				Message:   w.Message,
			},
		)
	}

	return result
}
