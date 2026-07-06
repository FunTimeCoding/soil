package model_context

import (
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func formatTime(t time.Time) string {
	return t.Format(timeLibrary.DateMinute)
}
