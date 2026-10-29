package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/service"
)

func New(
	v *service.Service,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	x, cancel := context.WithCancel(context.Background())

	return &Worker{
		service:  v,
		logger:   l,
		recovery: recovery.New(l, r),
		context:  x,
		cancel:   cancel,
	}
}
