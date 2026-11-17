package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsOutItsInventoryGuardFirst(t *testing.T) {
	s := base.New(t)
	s.UnreachableInventory()
	workers := s.Daemon().Workers()
	assert.Count(t, 3, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n == 0, guard)
	}
}

func TestDaemonCountsReapsOnItsRegistry(t *testing.T) {
	s := base.New(t)
	s.Daemon()
	assert.True(t, reaper.CountsOn(s.Registry))
}

func TestUnreachableInventoryIsLeftToGuard(t *testing.T) {
	s := base.New(t)
	s.UnreachableInventory()
	s.Start()
	r := s.WaitForRun()
	assert.String(t, "error", r.Status)
	assert.String(t, "inventory", r.Scope)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := base.New(t)
	s.UnreachableInventory()
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	s.WaitForRun()

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
