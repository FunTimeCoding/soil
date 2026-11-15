package jellyfin_tester

import (
	"github.com/funtimecoding/soil/pkg/jellyfin"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	address string,
) *jellyfin.Client {
	t.Helper()
	host, port := upstream_server.HostPort(t, address)
	result := jellyfin.New(host, port, "alfa-token")
	t.Cleanup(result.Close)

	return result
}
