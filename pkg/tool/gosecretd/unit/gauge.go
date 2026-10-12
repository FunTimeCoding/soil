package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func gauge(
	t *testing.T,
	y *prometheus.Registry,
	name string,
) float64 {
	t.Helper()
	families, e := y.Gather()
	assert.FatalOnError(t, e)

	for _, f := range families {
		if f.GetName() == name {
			return f.GetMetric()[0].GetGauge().GetValue()
		}
	}

	t.Fatalf("no gauge %s", name)

	return 0
}
