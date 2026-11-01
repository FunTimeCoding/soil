package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v *service.Service,
	interval time.Duration,
	l *logger.Logger,
	registry *prometheus.Registry,
	r face.Reporter,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		service:  v,
		interval: interval,
		logger:   l,
		reporter: r,
		metric:   metric.New(registry),
		recovery: recovery.New(l, r),
		context:  x,
		cancel:   cancel,
		stop:     make(chan struct{}),
	}
}
