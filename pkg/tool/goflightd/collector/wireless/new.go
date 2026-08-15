package wireless

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"time"
)

func New(
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
) *Collector {
	return &Collector{
		store:    s,
		logger:   l,
		recovery: recovery.New(l, r),
		interval: interval,
		stop:     make(chan struct{}),
		last:     make(map[string]string),
	}
}
