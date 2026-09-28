package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	s *store.Store,
	collectors []face.Collector,
	d *lease.Collector,
	n *client.ClientWithResponses,
	interval time.Duration,
	retention time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
	r soilFace.Reporter,
) *Worker {
	return &Worker{
		store:      s,
		collectors: collectors,
		lease:      d,
		netbox:     n,
		interval:   interval,
		retention:  retention,
		logger:     l,
		metric:     metric.New(y),
		recovery:   recovery.New(l, r),
	}
}
