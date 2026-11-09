package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/option"
	"testing"
)

func New(
	t *testing.T,
	o *option.Certificate,
	r *memory.Memory,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(o, logger.New(t.Context()), r)
	t.Cleanup(result.Close)

	return result
}
