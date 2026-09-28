package worker

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/lease"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/metric"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"time"
)

type Worker struct {
	store      *store.Store
	collectors []face.Collector
	lease      *lease.Collector
	netbox     *client.ClientWithResponses
	interval   time.Duration
	retention  time.Duration
	logger     *logger.Logger
	metric     *metric.Metric
	recovery   *recovery.Recovery
	cancel     context.CancelFunc
}
