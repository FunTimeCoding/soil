package daemon_tester

import (
	alertmanagerMock "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusMock "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/daemon"
	"testing"
)

func NewHealthy(
	t *testing.T,
	names ...string,
) *daemon.Daemon {
	t.Helper()
	alertmanagers := map[string]*alertmanagerMock.Client{}
	prometheuses := map[string]*prometheusMock.Client{}

	for _, n := range names {
		alertmanagers[n] = alertmanagerMock.New()
		prometheuses[n] = prometheusMock.New()
	}

	return New(t, alertmanagers, prometheuses)
}
