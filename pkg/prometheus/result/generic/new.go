package generic

import "time"

func New(
	typeName string,
	metric string,
	t time.Time,
	value string,
) *Result {
	return &Result{Type: typeName, Metric: metric, Time: t, Value: value}
}
