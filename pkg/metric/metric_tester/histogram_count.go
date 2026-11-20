package metric_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/client_golang/prometheus"
)

func (o *Tester) HistogramCount(
	r *prometheus.Registry,
	name string,
) uint64 {
	o.t.Helper()
	families, e := r.Gather()
	assert.FatalOnError(o.t, e)

	for _, f := range families {
		if f.GetName() != name {
			continue
		}

		for _, m := range f.GetMetric() {
			return m.GetHistogram().GetSampleCount()
		}
	}

	return 0
}
