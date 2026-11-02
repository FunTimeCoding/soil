package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"testing"
)

func New(t *testing.T) (*ansible_tester.Tester, *daemon.Daemon) {
	t.Helper()
	f := ansible_tester.NewFixture(t, nil)

	return f, Build(t, f)
}
