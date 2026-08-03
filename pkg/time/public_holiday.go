package time

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func PublicHoliday(t time.Time) bool {
	for _, d := range constant.PublicHolidays {
		if d.Year() == t.Year() &&
			d.Month() == t.Month() &&
			d.Day() == t.Day() {
			return true
		}
	}

	return false
}
