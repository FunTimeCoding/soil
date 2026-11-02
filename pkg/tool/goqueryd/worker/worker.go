package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/metric"
	"sync"
	"time"
)

type Worker struct {
	context  context.Context
	cancel   context.CancelFunc
	maintain func(context.Context) (int64, error)
	interval time.Duration
	logger   *logger.Logger
	reporter face.Reporter
	metric   *metric.Metric
	recovery *recovery.Recovery
	started  sync.Once
	done     chan struct{}
}
