package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Organization = "test"
	o.Host = "sentry.test"
	o.Source = func() face.SentryClient { return c }
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
