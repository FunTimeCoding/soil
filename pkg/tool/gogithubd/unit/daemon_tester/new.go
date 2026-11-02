package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/mock_source"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/option"
	"testing"
)

func New(
	t *testing.T,
	s *mock_source.Source,
	m *metric.Metric,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Owner = "alfa"
	o.Source = func() face.RunSource { return s }
	result := daemon.New(o, logger.New(t.Context()), memory.New(), m)
	t.Cleanup(result.Close)

	return result
}
