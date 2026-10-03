package holiday

import "time"

func NewFixed(
	m time.Month,
	day int,
) *Holiday {
	return &Holiday{month: m, day: day}
}
