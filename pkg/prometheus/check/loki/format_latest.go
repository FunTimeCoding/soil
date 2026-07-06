package loki

import "github.com/funtimecoding/soil/pkg/time"

func formatLatest(e *overview) string {
	if e.Latest.IsZero() {
		return ""
	}

	return e.Latest.Format(time.DateMinute)
}
