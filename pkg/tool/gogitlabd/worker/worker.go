package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Worker struct {
	client   *gitlab.Client
	interval time.Duration
	gauge    *prometheus.GaugeVec
	recovery *recovery.Recovery
	stop     chan struct{}
}
