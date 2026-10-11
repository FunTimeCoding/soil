package time

import "time"

func MonthsBetween(
	from time.Time,
	to time.Time,
) int {
	result := (to.Year()-from.Year())*12 + int(to.Month()) - int(from.Month())

	if to.Day() < from.Day() {
		result--
	}

	return result
}
