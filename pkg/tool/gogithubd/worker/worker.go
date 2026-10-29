package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Worker struct {
	service  face.Latest
	interval time.Duration
	gauge    *prometheus.GaugeVec
	recovery *recovery.Recovery
	stop     chan struct{}
}
