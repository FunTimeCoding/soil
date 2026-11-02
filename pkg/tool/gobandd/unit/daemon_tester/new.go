package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/band/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/option"
	"testing"
)

func New(
	t *testing.T,
	names ...string,
) (*daemon.Daemon, map[string]*mock_client.Client) {
	t.Helper()
	engines := make(map[string]*mock_client.Client)
	var instances []instance.Instance

	for _, name := range names {
		engines[name] = mock_client.New()
		instances = append(
			instances,
			instance.Instance{Name: name, Host: join.Empty(name, ".example")},
		)
	}

	o := option.New()
	o.Inventory = inventory.New[instance.Instance](instances...)
	o.Engine = func(i *instance.Instance) (face.BandClient, error) {
		return engines[i.Name], nil
	}
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result, engines
}
