package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/outcome"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	collect func(context.Context) ([]*outcome.Outcome, error),
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		context:  x,
		cancel:   cancel,
		collect:  collect,
		interval: interval,
		logger:   l,
		reporter: r,
		metric:   metric.New(y),
		recovery: recovery.New(l, r),
		done:     make(chan struct{}),
	}
}
