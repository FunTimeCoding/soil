package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/service"
	"time"
)

func New(
	v *service.Service,
	interval time.Duration,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	return &Worker{
		service:  v,
		interval: interval,
		logger:   l,
		reporter: r,
		recovery: recovery.New(l, r),
		notifier: notifier.New(),
		stop:     make(chan struct{}),
	}
}
