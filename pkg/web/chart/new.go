package chart

import "time"

func New(
	start time.Time,
	end time.Time,
) *Chart {
	return &Chart{start: start, end: end, maximum: 100}
}
