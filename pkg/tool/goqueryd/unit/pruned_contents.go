package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func prunedContents(r *prometheus.Registry) float64 {
	families, e := r.Gather()
	errors.PanicOnError(e)

	for _, f := range families {
		if f.GetName() == "goqueryd_pruned_contents_total" {
			return f.GetMetric()[0].GetCounter().GetValue()
		}
	}

	return 0
}
