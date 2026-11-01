package reconciler

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"time"
)

type Reconciler struct {
	service  *service.Service
	interval time.Duration
	logger   *logger.Logger
	reporter face.Reporter
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
}
