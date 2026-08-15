package server

import "time"

func timeRange(
	start *time.Time,
	end *time.Time,
) (time.Time, time.Time) {
	to := time.Now()

	if end != nil {
		to = *end
	}

	from := to.Add(-1 * time.Hour)

	if start != nil {
		from = *start
	}

	return from, to
}
