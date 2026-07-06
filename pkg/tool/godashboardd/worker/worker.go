package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"time"
)

type Worker struct {
	service  *service.Service
	interval time.Duration
	recovery *recovery.Recovery
	stop     chan struct{}
}
