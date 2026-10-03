package unit

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func holidaysBetween(
	start time.Time,
	end time.Time,
) []time.Time {
	var result []time.Time

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if library.PublicHoliday(d) {
			result = append(result, d)
		}
	}

	return result
}
