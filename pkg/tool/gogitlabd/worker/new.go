package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	refresh func(context.Context) []error,
	entries func() []*board_entry.Entry,
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
		[]string{constant.SourceLabel, "project", "ref", "status"},
	)
	y.MustRegister(g)
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		refresh:  refresh,
		entries:  entries,
		interval: interval,
		gauge:    g,
		reporter: r,
		logger:   l,
		recovery: recovery.New(l, r),
		context:  x,
		cancel:   cancel,
		stop:     make(chan struct{}),
		done:     make(chan struct{}),
	}
}
