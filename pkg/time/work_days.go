package time

import "time"

func WorkDays(
	start time.Time,
	end time.Time,
) int {
	var result int

	if start.After(end) {
		start, end = end, start
	}

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday {
			continue
		}

		if d.Weekday() == time.Sunday {
			continue
		}

		if PublicHoliday(d) {
			continue
		}

		result++
	}

	return result
}
