package dispatch_tester

import "time"

func At() time.Time {
	return time.Date(2026, 9, 8, 14, 32, 0, 0, time.Local)
}
