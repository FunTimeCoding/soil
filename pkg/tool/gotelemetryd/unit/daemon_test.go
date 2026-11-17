package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	b := base.New(t)
	before := goleak.IgnoreCurrent()
	b.Daemon().Close()
	goleak.VerifyNone(t, before)
}
