package watch

import (
	"github.com/funtimecoding/soil/pkg/openbao/status"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
)

type Watch struct {
	status    func() (*status.Status, error)
	sealed    prometheus.Gauge
	reachable prometheus.Gauge
	clock     *worker.Worker
}
