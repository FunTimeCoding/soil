package service_tester

import "time"

func Now() time.Time {
	return time.Date(2030, 1, 2, 3, 4, 5, 0, time.UTC)
}
