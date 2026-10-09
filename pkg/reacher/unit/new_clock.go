package unit

import "time"

func newClock() *clock {
	return &clock{now: time.Date(2026, 10, 6, 11, 23, 0, 0, time.UTC)}
}
