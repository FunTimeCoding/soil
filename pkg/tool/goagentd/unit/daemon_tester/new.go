package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/option"
	"io"
	"testing"
)

func New(
	t *testing.T,
	stream func(string, io.Writer, ...string) face.Stream,
) *daemon.Daemon {
	t.Helper()
	o := option.New()
	o.Workspace = t.TempDir()
	o.Stream = stream
	result := daemon.New(o, logger.New(t.Context()), memory.New())
	t.Cleanup(result.Close)

	return result
}
