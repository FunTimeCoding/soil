package time

import "time"

func Months(
	start time.Time,
	end time.Time,
) int {
	result := (end.Year()-start.Year())*12 + int(end.Month()) - int(start.Month())

	if end.Day() < start.Day() {
		result--
	}

	return result
}
