package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func histogramCount(
	t *testing.T,
	r *prometheus.Registry,
	name string,
) uint64 {
	t.Helper()
	families, e := r.Gather()
	assert.FatalOnError(t, e)

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
