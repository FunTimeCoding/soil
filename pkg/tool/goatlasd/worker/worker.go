package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/types/outcome"
	"sync"
	"time"
)

type Worker struct {
	collect  func(context.Context) ([]*outcome.Outcome, error)
	interval time.Duration
	logger   *logger.Logger
	reporter face.Reporter
	metric   *metric.Metric
	recovery *recovery.Recovery
	cancel   context.CancelFunc
	context  context.Context
	started  sync.Once
	done     chan struct{}
}
