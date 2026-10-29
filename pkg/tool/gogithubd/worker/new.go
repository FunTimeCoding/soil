package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	soil "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	v face.Latest,
	interval time.Duration,
	y *prometheus.Registry,
	l *logger.Logger,
	r soil.Reporter,
) *Worker {
	g := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "github_workflow_run_status",
			Help: "Latest completed workflow run per repo/workflow/branch (1 per entry)",
		},
		[]string{"owner", "repo", "workflow", "branch", "conclusion"},
	)
	y.MustRegister(g)

	return &Worker{
		service:  v,
		interval: interval,
		gauge:    g,
		recovery: recovery.New(l, r),
		stop:     make(chan struct{}),
	}
}
