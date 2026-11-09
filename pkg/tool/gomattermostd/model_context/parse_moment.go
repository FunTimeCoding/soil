package model_context

import (
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func parseMoment(s string) (time.Time, error) {
	t, e := time.ParseInLocation(
		timeConstant.DateMinute,
		s,
		time.Now().Location(),
	)

	if e != nil {
		return time.Parse(time.RFC3339, s)
	}

	return t, nil
}
