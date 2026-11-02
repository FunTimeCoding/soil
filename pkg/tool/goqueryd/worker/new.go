package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/metric"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	maintain func(context.Context) (int64, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		context:  x,
		cancel:   cancel,
		maintain: maintain,
		interval: interval,
		logger:   l,
		reporter: r,
		metric:   metric.New(y),
		recovery: recovery.New(l, r),
		done:     make(chan struct{}),
	}
}
