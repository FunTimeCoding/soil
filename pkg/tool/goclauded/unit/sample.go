package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/client_golang/prometheus"
	"maps"
	"testing"
)

func sample(
	t *testing.T,
	r *prometheus.Registry,
	name string,
	labels map[string]string,
) float64 {
	t.Helper()
	families, e := r.Gather()
	assert.FatalOnError(t, e)

	for _, f := range families {
		if f.GetName() != name {
			continue
		}

		for _, m := range f.GetMetric() {
			found := map[string]string{}

			for _, l := range m.GetLabel() {
				found[l.GetName()] = l.GetValue()
			}

			if maps.Equal(found, labels) {
				return m.GetCounter().GetValue() + m.GetGauge().GetValue()
			}
		}
	}

	return -1
}
