package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsOutItsInventoryGuardFirst(t *testing.T) {
	f := ansible_tester.NewFixture(t, nil)
	f.UnreachableInventory()
	workers := daemon_tester.Build(t, f).Workers()
	assert.Count(t, 3, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n == 0, guard)
	}
}

func TestDaemonCountsReapsOnItsRegistry(t *testing.T) {
	f := ansible_tester.NewFixture(t, nil)
	daemon_tester.Build(t, f)
	assert.True(t, reaper.CountsOn(f.Registry))
}

func TestUnreachableInventoryIsLeftToGuard(t *testing.T) {
	f := ansible_tester.NewFixture(t, nil)
	f.UnreachableInventory()
	daemon_tester.Start(daemon_tester.Build(t, f))
	r := f.WaitForRun()
	assert.String(t, "error", r.Status)
	assert.String(t, "inventory", r.Scope)
	assert.Count(t, 0, f.Reporter.Events())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	f := ansible_tester.NewFixture(t, nil)
	f.UnreachableInventory()
	before := goleak.IgnoreCurrent()
	d := daemon_tester.Build(t, f)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	f.WaitForRun()

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
