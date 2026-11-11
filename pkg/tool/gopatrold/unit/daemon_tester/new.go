package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/patrol/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/option"
	"testing"
)

func New(
	t *testing.T,
	clients map[string]*mock_client.Client,
	m *metric.Metric,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	o := option.New()

	for host := range clients {
		o.Hosts = append(o.Hosts, host)
	}

	o.Source = func(host string) face.PatrolClient { return clients[host] }
	result := daemon.New(o, logger.New(t.Context()), r, m)
	t.Cleanup(result.Close)

	return result
}
