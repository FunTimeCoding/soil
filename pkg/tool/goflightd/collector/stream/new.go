package stream

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
)

func New(
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	predicate string,
) *Collector {
	return &Collector{
		store:     s,
		logger:    l,
		recovery:  recovery.New(l, r),
		predicate: predicate,
		stop:      make(chan struct{}),
	}
}
