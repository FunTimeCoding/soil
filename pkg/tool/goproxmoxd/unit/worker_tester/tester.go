package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

type Tester struct {
	*service_tester.Tester
	Registry *prometheus.Registry
	worker   *worker.Worker
	t        *testing.T
}
