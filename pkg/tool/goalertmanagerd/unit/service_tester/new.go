package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/face"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusFace "github.com/funtimecoding/soil/pkg/prometheus/face"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/option"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/types/instance"
	"maps"
	"slices"
	"testing"
)

func New(
	t *testing.T,
	alertmanagers map[string]*mock_client.Client,
	prometheuses map[string]*prometheusMock.Client,
) *service.Service {
	t.Helper()
	var instances []instance.Instance

	for _, name := range slices.Sorted(maps.Keys(alertmanagers)) {
		i := instance.Instance{Name: name}
		i.Alertmanager.Host = join.Empty("alertmanager.", name, ".example")
		i.Prometheus.Host = join.Empty("prometheus.", name, ".example")
		instances = append(instances, i)
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](instances...)
	o.Alertmanager = func(i *instance.Instance) (face.Client, error) {
		return alertmanagers[i.Name], nil
	}
	o.Prometheus = func(i *instance.Instance) (prometheusFace.Client, error) {
		return prometheuses[i.Name], nil
	}
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
