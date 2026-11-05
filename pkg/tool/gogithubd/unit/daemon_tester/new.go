package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/github/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"testing"
)

func New(
	t *testing.T,
	c *mock_client.Client,
	m *metric.Metric,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Owner = "alfa"
	o.Source = func() face.GitHubClient { return c }
	result := daemon.New(o, logger.New(t.Context()), r, m)
	t.Cleanup(result.Close)

	return result
}
