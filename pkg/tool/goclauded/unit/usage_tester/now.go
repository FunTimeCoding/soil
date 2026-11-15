package usage_tester

import "time"

func Now() time.Time {
	return time.Date(2026, 11, 9, 12, 0, 0, 0, time.UTC)
}
