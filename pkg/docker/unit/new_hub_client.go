package unit

import (
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func newHubClient(
	t *testing.T,
	server string,
) *hub.Client {
	t.Helper()
	result := hub.NewWithBase(
		upstream_tester.Locator(t, server).Base(constant.BasePath),
	).WithBackoff(0)
	t.Cleanup(result.Close)

	return result
}
