package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
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
		recovery: recovery.New(l, r),
		stop:     make(chan struct{}),
	}
}
