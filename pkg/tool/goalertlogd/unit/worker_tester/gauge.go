package worker_tester

import "github.com/funtimecoding/soil/pkg/assert"

func (o *Tester) Gauge(name string) float64 {
	o.t.Helper()
	families, e := o.Registry.Gather()
	assert.FatalOnError(o.t, e)

	for _, f := range families {
		if f.GetName() == name {
			return f.GetMetric()[0].GetGauge().GetValue()
		}
	}

	o.t.Fatalf("no gauge %s", name)

	return 0
}
