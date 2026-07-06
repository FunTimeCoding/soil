package model_context

import (
	timeLibrary "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func parseSince(s string) (time.Time, error) {
	t, e := time.ParseInLocation(
		timeLibrary.DateMinute,
		s,
		time.Now().Location(),
	)

	if e != nil {
		return time.Parse(time.RFC3339, s)
	}

	return t, nil
}
