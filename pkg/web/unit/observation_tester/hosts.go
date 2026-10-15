package observation_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"slices"
)

func (o *Observation) Hosts() []string {
	var result []string

	for _, l := range labels(o.registry) {
		if h := l[constant.MetricHostLabel]; !slices.Contains(result, h) {
			result = append(result, h)
		}
	}

	return result
}
