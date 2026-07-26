package time

import (
	"github.com/funtimecoding/soil/pkg/time/constant"
	"time"
)

func FormatCompact(t time.Time) string {
	l := t.Local()
	n := time.Now()

	if l.Year() == n.Year() && l.YearDay() == n.YearDay() {
		return l.Format(constant.HourMinute)
	}

	return l.Format(constant.DateMinute)
}
