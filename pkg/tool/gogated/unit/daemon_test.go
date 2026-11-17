package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesOnce(t *testing.T) {
	d := base.New(t).Daemon()
	d.Close()
	d.Close()
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := base.New(t).Daemon()
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
	assert.Count(t, 1, base.New(t).Daemon().Workers())
}
