package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"testing"
)

func New(
	t *testing.T,
	f *mock_client.Client,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(
		store_tester.NewOption(t, f),
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(result.Close)

	return result
}
