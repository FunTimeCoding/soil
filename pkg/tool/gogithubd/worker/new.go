package worker

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	client *github.Client,
	owner string,
	interval time.Duration,
	y *prometheus.Registry,
	l *logger.Logger,
	r face.Reporter,
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
		client:   client,
		owner:    owner,
		interval: interval,
		gauge:    g,
		recovery: recovery.New(l, r),
		stop:     make(chan struct{}),
	}
}
