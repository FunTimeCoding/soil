package watcher

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
)

func New(
	rescan func(context.Context) []error,
	sync func(string, string),
	directories func() []string,
	harbor string,
	l *logger.Logger,
	r face.Reporter,
) *Watcher {
	x, cancel := context.WithCancel(context.Background())

	return &Watcher{
		context:     x,
		cancel:      cancel,
		rescan:      rescan,
		sync:        sync,
		directories: directories,
		recovery:    recovery.New(l, r),
		reporter:    r,
		harbor:      harbor,
		changed:     make(map[string]string),
		done:        make(chan struct{}),
	}
}
