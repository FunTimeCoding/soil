package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func Build(
	t *testing.T,
	f *ansible_tester.Tester,
) *daemon.Daemon {
	t.Helper()
	f.Registry = prometheus.NewRegistry()
	result := daemon.New(
		f.Option,
		logger.New(t.Context()),
		f.Reporter,
		f.Registry,
	)
	t.Cleanup(result.Close)

	return result
}
