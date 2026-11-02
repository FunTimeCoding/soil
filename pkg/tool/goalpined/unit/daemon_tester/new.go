package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goalpined/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"testing"
)

func New(
	t *testing.T,
	o *option.Alpine,
) *daemon.Daemon {
	t.Helper()

	return daemon.New(o)
}
