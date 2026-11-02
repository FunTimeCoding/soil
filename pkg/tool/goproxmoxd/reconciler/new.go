package reconciler

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"time"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	interval time.Duration,
	l *logger.Logger,
	r face.Reporter,
) *Reconciler {
	x, cancel := context.WithCancel(context.Background())

	return &Reconciler{
		service:   v,
		inventory: i,
		interval:  interval,
		logger:    l,
		reporter:  r,
		recovery:  recovery.New(l, r),
		context:   x,
		cancel:    cancel,
	}
}
