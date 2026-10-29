package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/metric"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"sync/atomic"
	"time"
)

type Worker struct {
	service   *service.Service
	notifier  face.EventNotifier
	logger    *logger.Logger
	reporter  face.Reporter
	recovery  *recovery.Recovery
	interval  time.Duration
	stop      chan struct{}
	closeOnce sync.Once
	lastPoll  atomic.Value
	registry  *prometheus.Registry
	metrics   *metric.Metric
}
