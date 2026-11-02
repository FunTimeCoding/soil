package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/postgres/face"
	"github.com/funtimecoding/soil/pkg/postgres/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/option"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service"
	"testing"
)

func New(
	t *testing.T,
	names ...string,
) (*service.Service, map[string]*mock_client.Client) {
	t.Helper()
	clients := make(map[string]*mock_client.Client)
	var instances []instance.Instance

	for _, name := range names {
		clients[name] = mock_client.New()
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

	return result, clients
}
