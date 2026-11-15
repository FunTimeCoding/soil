package clock_tester

import "time"

func NewAt(start time.Time) *Clock {
	return &Clock{now: start}
}
