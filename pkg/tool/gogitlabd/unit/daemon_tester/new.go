package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/inventory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/option"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	names ...string,
) (*daemon.Daemon, map[string]*mock_client.Client) {
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
	o.Client = func(i *instance.Instance) (face.GitLabClient, error) {
		return clients[i.Name], nil
	}
	result := daemon.New(
		o,
		logger.New(t.Context()),
		r,
		prometheus.NewRegistry(),
	)
	t.Cleanup(result.Close)

	return result, clients
}
