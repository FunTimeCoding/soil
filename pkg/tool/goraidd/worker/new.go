package worker

import (
	"context"
	"github.com/fsnotify/fsnotify"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	handle func(context.Context, fsnotify.Event) []error,
	paths []string,
	l *logger.Logger,
	r face.Reporter,
	prepare ...func(context.Context) []error,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		handle:   handle,
		paths:    paths,
		prepare:  prepare,
		logger:   l,
		reporter: r,
		recovery: recovery.New(l, r),
		context:  x,
		cancel:   cancel,
	}
}
