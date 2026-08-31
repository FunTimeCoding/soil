package usage_result

import "time"

type Result struct {
	FiveHourPercent int
	FiveHourReset   time.Time
	SevenDayPercent int
	SevenDayReset   time.Time
	FablePercent    int
	FableReset      string
	LastUpdated     time.Time
}
