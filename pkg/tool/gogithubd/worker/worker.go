package worker

import (
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/types/run_key"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

type Worker struct {
	latest func() (map[run_key.Key]*run.Run, error)
	owner  string
	gauge  *prometheus.GaugeVec
	clock  *worker.Worker
}
