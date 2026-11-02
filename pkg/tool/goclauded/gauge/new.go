package gauge

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	findings func() ([]*finding.Finding, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Gauge {
	g := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "goclauded_findings",
			Help: "Inconsistencies found in the coordination state, by kind.",
		},
		[]string{constant.Kind},
	)
	y.MustRegister(g)
	x, cancel := context.WithCancel(context.Background())

	return &Gauge{
		context:  x,
		cancel:   cancel,
		findings: findings,
		interval: interval,
		reporter: r,
		recovery: recovery.New(l, r),
		gauge:    g,
		done:     make(chan struct{}),
	}
}
