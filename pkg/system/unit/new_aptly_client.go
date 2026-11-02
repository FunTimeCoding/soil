package unit

import (
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newAptlyClient(
	t *testing.T,
	server string,
) *aptly.Client {
	t.Helper()
	host, port := upstream_tester.HostPort(t, server)
	result := aptly.New(host, port, true, "alfa", "bravo")
	t.Cleanup(result.Close)

	return result
}
