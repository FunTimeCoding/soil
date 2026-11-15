package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/linkace/mock_client"
	"testing"
)

func ServedPort(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) int {
	t.Helper()

	return served(t, c, r).Port
}
