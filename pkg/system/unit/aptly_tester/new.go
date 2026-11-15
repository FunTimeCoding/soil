package aptly_tester

import (
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *aptly.Client {
	t.Helper()
	host, port := upstream_server.HostPort(t, address)
	result := aptly.New(host, port, true, "alfa", "bravo")
	t.Cleanup(result.Close)

	return result
}
