package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"testing"
)

func Build(
	t *testing.T,
	f *ansible_tester.Tester,
) *daemon.Daemon {
	t.Helper()
	result := daemon.New(f.Option, logger.New(t.Context()), f.Reporter)
	t.Cleanup(result.Close)

	return result
}
