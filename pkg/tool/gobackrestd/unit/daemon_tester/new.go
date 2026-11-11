package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/backrest/face"
	"github.com/funtimecoding/soil/pkg/backrest/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/option"
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
	o.Source = func() face.BackrestClient { return c }
	result := daemon.New(o, logger.New(t.Context()), r, m)
	t.Cleanup(result.Close)

	return result
}
