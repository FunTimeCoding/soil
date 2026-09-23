package gauge

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func value(
	t *testing.T,
	r *prometheus.Registry,
	kind string,
) float64 {
	t.Helper()
	families, e := r.Gather()
	assert.FatalOnError(t, e)

	for _, f := range families {
		if f.GetName() != "goclauded_findings" {
			continue
		}

		for _, m := range f.GetMetric() {
			for _, l := range m.GetLabel() {
				if l.GetName() == "kind" && l.GetValue() == kind {
					return m.GetGauge().GetValue()
				}
			}
		}
	}

	return -1
}
