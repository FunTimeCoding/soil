package hub_tester

import (
	"github.com/funtimecoding/soil/pkg/docker/constant"
	"github.com/funtimecoding/soil/pkg/docker/hub"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	server string,
) *hub.Client {
	t.Helper()
	result := hub.NewWithBase(
		locator.NewOrigin(server).Base(constant.BasePath),
	).WithBackoff(0)
	t.Cleanup(result.Close)

	return result
}
