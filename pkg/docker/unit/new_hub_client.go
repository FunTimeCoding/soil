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

	return hub.NewWithBase(
		upstream_tester.Locator(t, server).Base(constant.BasePath),
	)
}
