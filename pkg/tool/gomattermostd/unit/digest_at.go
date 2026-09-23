package unit

import "time"

func digestAt(
	hour int,
	minute int,
) time.Time {
	return time.Date(2026, 9, 8, hour, minute, 0, 0, time.UTC)
}
