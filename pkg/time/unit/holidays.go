package unit

import (
	soilTime "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func holidays(
	start time.Time,
	end time.Time,
) []time.Time {
	var result []time.Time

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if soilTime.PublicHoliday(d) {
			result = append(result, d)
		}
	}

	return result
}
