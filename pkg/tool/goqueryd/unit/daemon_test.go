package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesWhatItBuiltOnce(t *testing.T) {
	q := service_tester.New(t)
	d := q.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, q.Embedder.Closed())
	assert.Integer(t, 1, q.Reranker.Closed)
}

func TestDaemonHandsOutItsGuardThenItsClock(t *testing.T) {
	assert.Count(t, 2, service_tester.New(t).Daemon().Workers())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	q := service_tester.New(t)
	before := goleak.IgnoreCurrent()
	d := q.Daemon()
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
