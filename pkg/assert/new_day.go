package assert

import "time"

func NewDay(day int) time.Time {
	return time.Date(1970, 1, day, 0, 0, 0, 0, time.UTC)
}
