package observation_tester

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func New() *Observation {
	h := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{Name: "alfa_seconds"},
		[]string{
			constant.MetricHostLabel,
			constant.MetricCodeLabel,
			constant.MetricMethodLabel,
		},
	)
	r := prometheus.NewRegistry()
	r.MustRegister(h)

	return &Observation{Histogram: h, registry: r}
}
