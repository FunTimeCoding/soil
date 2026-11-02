package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/daemon_tester"
	"testing"
)

func healthyDaemon(
	t *testing.T,
	names ...string,
) *daemon.Daemon {
	t.Helper()
	alertmanagers := map[string]*mock_client.Client{}
	prometheuses := map[string]*prometheusMock.Client{}

	for _, n := range names {
		alertmanagers[n] = mock_client.New()
		prometheuses[n] = prometheusMock.New()
	}

	return daemon_tester.New(t, alertmanagers, prometheuses)
}
