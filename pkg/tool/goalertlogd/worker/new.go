package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	a AlertSource,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
	retention time.Duration,
	y *prometheus.Registry,
) *Worker {
	p := &Worker{
		client:    a,
		store:     s,
		logger:    l,
		recovery:  recovery.New(l, r),
		interval:  interval,
		retention: retention,
		stop:      make(chan struct{}),
		registry:  y,
	}

	if y != nil {
		p.metrics = newMetrics(y)
	}

	return p
}
