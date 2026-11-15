package coverage_tester

import "time"

func Now() time.Time {
	return time.Date(2026, 8, 22, 12, 0, 0, 0, time.UTC)
}
