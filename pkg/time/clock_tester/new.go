package clock_tester

import "time"

func New() *Clock {
	return NewAt(time.Date(2026, 10, 6, 11, 23, 0, 0, time.UTC))
}
