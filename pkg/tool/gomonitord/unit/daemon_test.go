package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	before := goleak.IgnoreCurrent()
	s.Daemon().Close()
	goleak.VerifyNone(t, before)
}
