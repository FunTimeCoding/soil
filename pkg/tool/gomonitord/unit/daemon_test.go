package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	daemon_tester.New(t).Close()
	goleak.VerifyNone(t, before)
}
