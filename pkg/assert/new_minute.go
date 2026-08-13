package assert

import "time"

func NewMinute(minute int) time.Time {
	return time.Date(2000, 1, 1, 0, minute, 0, 0, time.UTC)
}
