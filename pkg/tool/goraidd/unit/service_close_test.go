package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/worker"
	soilWorker "github.com/funtimecoding/soil/pkg/worker"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesWhatItBuilt(t *testing.T) {
	p := service_tester.NewParser(t)
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(t, t.TempDir(), t.TempDir(), p),
		memory.New(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, p.Closed)
}

func TestDaemonHandsOutItsGuardWatcherAndCleanupClock(t *testing.T) {
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			t,
			t.TempDir(),
			t.TempDir(),
			service_tester.NewParser(t),
		),
		memory.New(),
	).Workers()
	assert.Count(t, 3, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
	_, watcher := workers[1].(*worker.Worker)
	assert.True(t, watcher)
	_, clock := workers[2].(*soilWorker.Worker)
	assert.True(t, clock)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	p := service_tester.NewParser(t)
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(t, t.TempDir(), t.TempDir(), p),
		memory.New(),
	)
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
