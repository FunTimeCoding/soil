package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Worker struct {
	client   *github.Client
	owner    string
	interval time.Duration
	gauge    *prometheus.GaugeVec
	recovery *recovery.Recovery
	stop     chan struct{}
}
