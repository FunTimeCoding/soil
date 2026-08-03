package day

import "time"

func New(
	year int,
	m time.Month,
	day int,
) time.Time {
	return time.Date(year, m, day, 0, 0, 0, 0, time.UTC)
}
