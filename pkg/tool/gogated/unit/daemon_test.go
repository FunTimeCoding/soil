package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesOnce(t *testing.T) {
	d := daemon.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	d.Close()
	d.Close()
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestDaemonWithoutDirectoryHandsOutOnlyItsClock(t *testing.T) {
	d := daemon.New(base.NewOption(t), logger.New(t.Context()), memory.New())
	t.Cleanup(d.Close)
	assert.Count(t, 1, d.Workers())
}
