package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"testing"
)

func newDaemon(
	t *testing.T,
	c face.NetBoxClient,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(
		service_tester.NewOption(t, c),
		logger.New(t.Context()),
		memory.New(),
		metric.New(),
	)
	t.Cleanup(result.Close)

	return result
}
