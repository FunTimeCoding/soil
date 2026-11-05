package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/sublime/face"
	"github.com/funtimecoding/soil/pkg/sublime/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Source = func() face.SublimeClient { return c }
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
