package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	proxmox "github.com/funtimecoding/soil/pkg/proxmox/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/collector"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	client func(string) (proxmox.ProxmoxClient, error),
	i *instance.Inventory,
	interval time.Duration,
	registry *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
) *Worker {
	return &Worker{
		client:    client,
		inventory: i,
		interval:  interval,
		collector: collector.New(registry),
		log:       l,
		recovery:  recovery.New(l, r),
		notifier:  notifier.New(),
		reacher:   reacher.New(r),
		stop:      make(chan struct{}),
		done:      make(chan struct{}),
	}
}
