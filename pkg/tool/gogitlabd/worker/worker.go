package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type Worker struct {
	refresh  func(context.Context) []error
	entries  func() []*board_entry.Entry
	interval time.Duration
	gauge    *prometheus.GaugeVec
	reporter face.Reporter
	logger   *logger.Logger
	recovery *recovery.Recovery
	context  context.Context
	cancel   context.CancelFunc
	stop     chan struct{}
	done     chan struct{}
	started  sync.Once
	stopped  sync.Once
}
