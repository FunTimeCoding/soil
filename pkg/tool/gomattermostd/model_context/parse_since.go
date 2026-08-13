package model_context

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func parseSince(s string) (time.Time, error) {
	t, e := time.ParseInLocation(constant.DateMinute, s, time.Now().Location())

	if e != nil {
		return time.Parse(time.RFC3339, s)
	}

	return t, nil
}
