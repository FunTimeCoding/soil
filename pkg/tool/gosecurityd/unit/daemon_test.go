package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/unit/base"
	"github.com/funtimecoding/soil/pkg/worker"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon.New(base.New(t).Option(), logger.New(t.Context()), memory.New())
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

func TestCloseStopsStartedWorkers(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon.New(base.New(t).Option(), logger.New(t.Context()), memory.New())

	for _, w := range d.Workers() {
		w.Start()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestWorkersHandGuardsBeforeClock(t *testing.T) {
	workers := base.New(t).Daemon().Workers()
	assert.Count(t, 4, workers)

	for _, w := range workers[:3] {
		_, guard := w.(*peer.Guard)
		assert.True(t, guard)
	}

	_, last := workers[3].(*worker.Worker)
	assert.True(t, last)
}

func TestCloseClosesEveryClient(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.Integer(t, 1, s.NationalDatabase.Closed)
	assert.Integer(t, 1, s.SecurityTracker.Closed)
	assert.Integer(t, 1, s.ExtendedTracker.Closed)
}

func TestClockLoadsBothTrackersAtStart(t *testing.T) {
	s := base.New(t)
	s.Started()
	assert.Eventually(
		t,
		func() bool {
			return s.SecurityTracker.Calls() > 0 &&
				s.ExtendedTracker.Calls() > 0
		},
	)
}
