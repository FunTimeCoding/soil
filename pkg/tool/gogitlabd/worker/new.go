package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v *service.Service,
	interval time.Duration,
	y *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	g := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gitlab_pipeline_status",
			Help: "Latest pipeline status per project and branch (1 per entry)",
		},
		[]string{"project", "ref", "status"},
	)
	y.MustRegister(g)

	return &Worker{
		service:  v,
		interval: interval,
		gauge:    g,
		recovery: recovery.New(l, r),
		notifier: notifier.New(),
		stop:     make(chan struct{}),
	}
}
