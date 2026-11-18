package battery_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/daemon"
)

type Tester struct {
	Metric   *metric.Metric
	Reporter *memory.Memory
	Daemon   *daemon.Daemon
}
