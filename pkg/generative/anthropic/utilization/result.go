package utilization

import "time"

type Result struct {
	SessionPercent int
	SessionReset   time.Time
	SessionSeen    bool
	WeeklyPercent  int
	WeeklyReset    time.Time
	FablePercent   int
	FableReset     time.Time
	FableSeen      bool
}
