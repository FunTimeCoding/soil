package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"sync"
	"time"
)

type Worker struct {
	service   *service.Service
	interval  time.Duration
	recovery  *recovery.Recovery
	reporter  face.Reporter
	logger    *logger.Logger
	stop      chan struct{}
	closeOnce sync.Once
}
