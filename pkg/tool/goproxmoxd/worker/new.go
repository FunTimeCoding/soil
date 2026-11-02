package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/sample"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	refresh func() []*sample.Sample,
	interval time.Duration,
	registry *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	return &Worker{
		refresh:   refresh,
		interval:  interval,
		collector: collector.New(registry),
		log:       l,
		reporter:  r,
		recovery:  recovery.New(l, r),
		stop:      make(chan struct{}),
		done:      make(chan struct{}),
	}
}
