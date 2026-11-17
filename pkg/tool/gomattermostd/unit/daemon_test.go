package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonWithoutOptionsRunsOnlyItsGuard(t *testing.T) {
	workers := service_tester.New(t).Daemon().Workers()
	assert.Count(t, 1, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
}

func TestGoclaudedAddsRosterGuardWatcherAndTwoClocks(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	workers := service_tester.New(
		t,
		service_tester.WithConnector(k),
	).Daemon().Workers()
	assert.Count(t, 5, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestMonitorAddsItsClock(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	workers := service_tester.New(
		t,
		service_tester.WithConnector(k),
		service_tester.WithMonitor(),
	).Daemon().Workers()
	assert.Count(t, 6, workers)
}

func TestDaemonClosesMattermostOnce(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	d := o.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, o.Mattermost.Closed)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	o.Mattermost.Failure = connection.New(
		errorsConstant.Unreachable,
		"alfa.test",
		"",
		errorsConstant.Refused,
	)
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
	d.Close()
	goleak.VerifyNone(t, before)
}
