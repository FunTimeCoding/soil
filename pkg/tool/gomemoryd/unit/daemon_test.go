package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/worker"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesTheQueryClientOnce(t *testing.T) {
	q := mock_client.New()
	t.Cleanup(q.Close)
	d := daemon.New(daemonOption(t, q), logger.New(t.Context()), memory.New())
	d.Close()
	d.Close()
	assert.Integer(t, 1, q.Closed)
}

func TestTheDaemonHandsItsGuardThenItsReconcileClockToTheLifecycle(
	t *testing.T,
) {
	d := daemon.New(
		daemonOption(t, mock_client.New()),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(d.Close)
	workers := d.Workers()
	assert.Count(t, 2, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
	_, clock := workers[1].(*worker.Worker)
	assert.True(t, clock)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		daemonOption(t, mock_client.New()),
		logger.New(t.Context()),
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
