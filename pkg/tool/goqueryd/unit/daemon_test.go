package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonClosesWhatItBuiltOnce(t *testing.T) {
	m := newEmbedder(t)
	a := newReranker(t)
	d := daemon.New(
		service_tester.NewOption(t, m, a),
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, m.Closed())
	assert.Integer(t, 1, a.Closed)
}

func TestDaemonHandsOutItsGuardThenItsClock(t *testing.T) {
	d := daemon.New(
		service_tester.NewOption(t, newEmbedder(t), newReranker(t)),
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	assert.Count(t, 2, d.Workers())
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.NewOption(t, newEmbedder(t), newReranker(t))
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
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
