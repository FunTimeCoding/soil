package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/metric"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v *service.Service,
	n face.EventNotifier,
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
	y *prometheus.Registry,
) *Worker {
	result := &Worker{
		service:  v,
		notifier: n,
		logger:   l,
		reporter: r,
		recovery: recovery.New(l, r),
		interval: interval,
		stop:     make(chan struct{}),
		registry: y,
	}

	if y != nil {
		result.metrics = metric.New(y)
	}

	return result
}
