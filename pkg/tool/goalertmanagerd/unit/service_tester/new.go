package service_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
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

	return pooled(
		t,
		instances,
		func(i *instance.Instance) face.AlertmanagerClient {
			return alertmanagers[i.Name]
		},
		func(i *instance.Instance) face.PrometheusClient {
			return prometheuses[i.Name]
		},
	)
}
