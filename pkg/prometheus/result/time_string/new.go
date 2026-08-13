package time_string

import "time"

func New(
	metric string,
	t time.Time,
	value string,
) *Result {
	return &Result{Metric: metric, Time: t, Value: value}
}
