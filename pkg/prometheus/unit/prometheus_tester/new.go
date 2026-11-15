package prometheus_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *prometheus.Client {
	t.Helper()
	host, port := upstream_server.HostPort(t, address)
	result := prometheus.New(host, port, false, "", "")
	t.Cleanup(result.Close)

	return result
}
