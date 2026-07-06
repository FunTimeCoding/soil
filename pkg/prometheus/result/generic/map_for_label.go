package generic

import "github.com/funtimecoding/soil/pkg/prometheus/result/metric"

func MapForLabel(
	r []*Result,
	label string,
) map[string][]*Result {
	result := make(map[string][]*Result)

	for _, e := range r {
		value := metric.New(e.Metric).Label(label)
		result[value] = append(result[value], e)
	}

	return result
}
