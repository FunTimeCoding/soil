package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	proxFace "github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v proxFace.Service,
	interval time.Duration,
	registry *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	return &Worker{
		service:   v,
		interval:  interval,
		collector: collector.New(registry),
		log:       l,
		recovery:  recovery.New(l, r),
		notifier:  notifier.New(),
		stop:      make(chan struct{}),
	}
}
