package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Host = "habitica.test"
	o.Source = func() face.HabiticaClient { return c }
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
