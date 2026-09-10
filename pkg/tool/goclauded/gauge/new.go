package gauge

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	s *service.Service,
	l *logger.Logger,
	r *prometheus.Registry,
) *Gauge {
	findings := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goclauded_findings",
			Help: "Inconsistencies found in the coordination state, by kind.",
		},
		[]string{constant.Kind},
	)
	r.MustRegister(findings)

	return &Gauge{service: s, logger: l, findings: findings}
}
