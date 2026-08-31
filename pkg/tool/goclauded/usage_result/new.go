package usage_result

import "time"

func New(
	fiveHourPercent int,
	fiveHourReset time.Time,
	sevenDayPercent int,
	sevenDayReset time.Time,
	fablePercent int,
	fableReset string,
	lastUpdated time.Time,
) *Result {
	return &Result{
		FiveHourPercent: fiveHourPercent,
		FiveHourReset:   fiveHourReset,
		SevenDayPercent: sevenDayPercent,
		SevenDayReset:   sevenDayReset,
		FablePercent:    fablePercent,
		FableReset:      fableReset,
		LastUpdated:     lastUpdated,
	}
}
