package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/metric"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v *service.Service,
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r face.Reporter,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		service:  v,
		interval: interval,
		logger:   l,
		reporter: r,
		metric:   metric.New(y),
		recovery: recovery.New(l, r),
		context:  x,
		cancel:   cancel,
	}
}
