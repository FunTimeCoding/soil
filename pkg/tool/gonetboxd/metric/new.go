package metric

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func New(r *prometheus.Registry) *Metric {
	result := &Metric{
		status: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: constant.StatusMetric,
				Help: constant.StatusHelp,
			},
			[]string{
				constant.KindLabel,
				constant.NameLabel,
				constant.StatusLabel,
			},
		),
	}
	r.MustRegister(result.status)

	return result
}
