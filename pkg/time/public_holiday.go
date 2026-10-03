package time

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func PublicHoliday(t time.Time) bool {
	for _, h := range constant.PublicHolidays {
		if h.On(t) {
			return true
		}
	}

	return false
}
