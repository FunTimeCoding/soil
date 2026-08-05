package rate_snapshot

import "time"

func New(
	fiveHourPercent int,
	sevenDayPercent int,
	fiveHourReset time.Time,
	sevenDayReset time.Time,
	createdAt time.Time,
) *Snapshot {
	return &Snapshot{
		FiveHourPercent: fiveHourPercent,
		SevenDayPercent: sevenDayPercent,
		FiveHourReset:   fiveHourReset,
		SevenDayReset:   sevenDayReset,
		CreatedAt:       createdAt,
	}
}
