package window

import "time"

func New(
	start time.Time,
	end time.Time,
) *Window {
	return &Window{start: &start, end: &end}
}
