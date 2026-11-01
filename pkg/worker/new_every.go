package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"time"
)

func NewEvery(
	tick func() []error,
	l *logger.Logger,
	r face.Reporter,
	interval time.Duration,
) *Worker {
	return New(func(context.Context) []error { return tick() }, l, r, interval)
}
