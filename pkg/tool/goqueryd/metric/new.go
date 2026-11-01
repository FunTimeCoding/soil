package metric

import "github.com/prometheus/client_golang/prometheus"

func New(r *prometheus.Registry) *Metric {
	m := &Metric{
		pruned: prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "goqueryd_pruned_contents_total",
				Help: "Contents swept because no document holds them",
			},
		),
	}
	r.MustRegister(m.pruned)

	return m
}
