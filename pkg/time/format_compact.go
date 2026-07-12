package time

import "time"

func FormatCompact(t time.Time) string {
	l := t.Local()
	n := time.Now()

	if l.Year() == n.Year() && l.YearDay() == n.YearDay() {
		return l.Format(HourMinute)
	}

	return l.Format(DateMinute)
}
