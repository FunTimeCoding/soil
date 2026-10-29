package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func New(
	tick func(context.Context) []error,
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		tick:     tick,
		logger:   l,
		reporter: r,
		recovery: recovery.New(l, r),
		interval: interval,
		context:  x,
		cancel:   cancel,
	}
}
