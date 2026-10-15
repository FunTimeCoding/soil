package metric

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/metric/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/prometheus/client_golang/prometheus"
)

func (m *Metric) Upstream(t *identity.Tool) prometheus.ObserverVec {
	m.once.Do(
		func() {
			m.upstream = prometheus.NewHistogramVec(
				prometheus.HistogramOpts{
					Name:    constant.UpstreamMetric,
					Help:    constant.UpstreamHelp,
					Buckets: prometheus.ExponentialBuckets(0.1, 2, 10),
				},
				[]string{
					constant.IdentityLabel,
					webConstant.MetricHostLabel,
					webConstant.MetricCodeLabel,
					webConstant.MetricMethodLabel,
				},
			)
			m.registry.MustRegister(m.upstream)
		},
	)

	return m.upstream.MustCurryWith(
		prometheus.Labels{constant.IdentityLabel: t.Name()},
	)
}
