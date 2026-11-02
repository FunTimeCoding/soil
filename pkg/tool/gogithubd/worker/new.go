package worker

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	latest func() (map[run_key.Key]*run.Run, error),
	owner string,
	interval time.Duration,
	l *logger.Logger,
	y *prometheus.Registry,
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
	result := &Worker{latest: latest, owner: owner, gauge: g}
	result.clock = worker.NewEvery(result.tick, l, r, interval)

	return result
}
