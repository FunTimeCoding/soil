package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonHandsItsMemoryGuardToLifecycle(t *testing.T) {
	d, _ := daemon_tester.New(t, inventory.New())
	workers := d.Workers()
	assert.Count(t, 1, workers)
	_, okay := workers[0].(*peer.Guard)
	assert.True(t, okay)
}

func TestClosingDaemonClosesMemoryClientOnce(t *testing.T) {
	d, m := daemon_tester.New(t, inventory.New())
	d.Close()
	d.Close()
	assert.Integer(t, 1, m.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d, _ := daemon_tester.New(t, inventory.New())
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
