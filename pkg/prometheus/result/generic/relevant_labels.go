package generic

import (
	"github.com/funtimecoding/soil/pkg/prometheus/result/metric"
	"github.com/funtimecoding/soil/pkg/strings"
)

func RelevantLabels(
	v []*Result,
	ignore []string,
) []string {
	var m []*metric.Metric

	for _, e := range v {
		m = append(m, metric.New(e.Metric))
	}

	return SortLabels(
		strings.RemoveFromList(metric.DifferingLabels(m), ignore),
	)
}
