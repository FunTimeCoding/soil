package observation_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func labels(r *prometheus.Registry) []map[string]string {
	families, e := r.Gather()
	errors.PanicOnError(e)
	var result []map[string]string

	for _, f := range families {
		for _, m := range f.GetMetric() {
			l := map[string]string{}

			for _, v := range m.GetLabel() {
				l[v.GetName()] = v.GetValue()
			}

			result = append(result, l)
		}
	}

	return result
}
