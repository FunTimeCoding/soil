package gauge

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type Gauge struct {
	context  context.Context
	cancel   context.CancelFunc
	findings func() ([]*finding.Finding, error)
	interval time.Duration
	reporter face.Reporter
	recovery *recovery.Recovery
	gauge    *prometheus.GaugeVec
	started  sync.Once
	done     chan struct{}
}
