package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"time"
)

type Worker struct {
	service  *service.Service
	interval time.Duration
	logger   *logger.Logger
	reporter face.Reporter
	metric   *metric.Metric
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
	stop     chan struct{}
}
