package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/time/time_of_day"
	"sync"
	"time"
)

type Worker struct {
	tick     func(context.Context) []error
	logger   *logger.Logger
	reporter face.Reporter
	recovery *recovery.Recovery
	interval time.Duration
	at       *time_of_day.Time
	once     bool
	context  context.Context
	cancel   context.CancelFunc
	started  sync.Once
	finished chan struct{}
}
