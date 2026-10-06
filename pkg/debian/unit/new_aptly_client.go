package unit

import (
	"github.com/funtimecoding/soil/pkg/debian/aptly"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newAptlyClient(
	t *testing.T,
	server string,
) *aptly.Client {
	t.Helper()
	host, port := upstream_tester.HostPort(t, server)

	return aptly.New(host, port, true, "alfa", "bravo")
}
