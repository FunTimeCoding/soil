package unit

import (
	"github.com/funtimecoding/soil/pkg/raid_parser"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newClient(
	t *testing.T,
	server string,
) *raid_parser.Client {
	t.Helper()

	return raid_parser.New(upstream_tester.Locator(t, server), "alfa-token")
}
