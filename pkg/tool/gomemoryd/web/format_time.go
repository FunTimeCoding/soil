package web

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func formatTime(raw string) string {
	t, e := time.Parse(time.RFC3339, raw)

	if e != nil {
		return raw
	}

	return library.FormatCompact(t)
}
