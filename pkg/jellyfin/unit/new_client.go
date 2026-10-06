package unit

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/basic"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
) *basic.Client {
	t.Helper()

	return basic.New(upstream_tester.Locator(t, server), "alfa-token")
}
