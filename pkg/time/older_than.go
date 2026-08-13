package time

import "time"

func OlderThan(
	t time.Time,
	minutes int,
) bool {
	return t.Before(time.Now().Add(time.Duration(minutes*-1) * time.Minute))
}
