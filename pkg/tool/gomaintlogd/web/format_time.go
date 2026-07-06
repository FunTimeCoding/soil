package web

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format(library.DateMinute)
}
