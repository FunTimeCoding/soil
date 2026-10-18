package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"time"
)

type Worker struct {
	service  *service.Service
	interval time.Duration
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
	stop     chan struct{}
}
