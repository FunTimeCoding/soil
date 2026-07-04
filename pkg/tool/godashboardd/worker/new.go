package worker

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/service"
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
