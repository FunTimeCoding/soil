package battery_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobatteryd/option"
	"testing"
)

func New(
	t *testing.T,
	instances ...*instance.Instance,
) *Tester {
	t.Helper()
	var values []instance.Instance

	for _, i := range instances {
		values = append(values, *i)
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](values...)
	result := &Tester{Metric: metric.New(), Reporter: memory.New()}
	result.Daemon = daemon.New(
		o,
		logger.New(t.Context()),
		result.Reporter,
		result.Metric,
	)
	t.Cleanup(result.Daemon.Close)

	return result
}
