package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := base.NewOption(t)
	before := goleak.IgnoreCurrent()
	daemon.New(o, logger.New(t.Context()), memory.New()).Close()
	goleak.VerifyNone(t, before)
}
