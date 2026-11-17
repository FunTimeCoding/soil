package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/worker"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesQueryClientOnce(t *testing.T) {
	o := service_tester.New(t)
	d := o.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, o.Indexer.Closed)
}

func TestDaemonHandsItsGuardThenItsReconcileClockToLifecycle(t *testing.T) {
	workers := service_tester.New(t).Daemon().Workers()
	assert.Count(t, 2, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
	_, clock := workers[1].(*worker.Worker)
	assert.True(t, clock)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.New(t)
	before := goleak.IgnoreCurrent()
	d := o.Daemon()
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
