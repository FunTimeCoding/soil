package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/face"
	"github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/option"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/service"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/types/instance"
	"maps"
	"slices"
	"testing"
)

func New(
	t *testing.T,
	clients map[string]*mock_client.Client,
) *service.Service {
	t.Helper()
	var instances []instance.Instance

	for _, name := range slices.Sorted(maps.Keys(clients)) {
		instances = append(
			instances,
			instance.Instance{Name: name, Host: join.Empty(name, ".example")},
		)
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](instances...)
	o.Client = func(i *instance.Instance) (face.Client, error) {
		return clients[i.Name], nil
	}
	result := service.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
