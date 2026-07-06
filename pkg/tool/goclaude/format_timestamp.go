package goclaude

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func formatTimestamp(ts string) string {
	t, e := time.Parse(time.RFC3339Nano, ts)

	if e != nil {
		t, e = time.Parse("2006-01-02T15:04:05.000Z", ts)
	}

	if e != nil {
		if len(ts) > 16 {
			return ts[:16]
		}

		return ts
	}

	return t.Local().Format(library.DateMinute)
}
