package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/prometheus/client_golang/prometheus"
	"sync/atomic"
	"time"
)

type Worker struct {
	client    AlertSource
	store     *store.Store
	logger    *logger.Logger
	recovery  *recovery.Recovery
	interval  time.Duration
	retention time.Duration
	stop      chan struct{}
	lastPoll  atomic.Value
	registry  *prometheus.Registry
	metrics   *metrics
}
