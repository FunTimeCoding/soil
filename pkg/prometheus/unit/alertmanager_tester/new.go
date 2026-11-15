package alertmanager_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	address string,
	o ...alertmanager.Option,
) *alertmanager.Client {
	t.Helper()
	host, port := upstream_server.HostPort(t, address)
	result := alertmanager.New(host, port, false, "", "", o...)
	t.Cleanup(result.Close)

	return result
}
