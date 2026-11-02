package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesBothPeerClientsOnce(t *testing.T) {
	o, m, q := newDaemonOption(t)
	defer m.Close()
	defer q.Close()
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, m.Closed)
	assert.Integer(t, 1, q.Closed)
}

func TestTheDaemonHandsOutItsGuardsFirst(t *testing.T) {
	o, _, _ := newDaemonOption(t)
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	workers := d.Workers()
	assert.Count(t, 11, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestTheDaemonClearsStaleBindingsBeforeItServes(t *testing.T) {
	o, _, _ := newDaemonOption(t)
	s := store_tester.NewOnFile(t, o.LitePath)
	r := s.EnsureSession("session-1")
	s.BindModelContextSession(r.Callsign, "mcp-session-abc")
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	assert.String(t, "", s.GetSession("session-1").ModelContextSession)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	o, _, _ := newDaemonOption(t)
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	workers := append(d.Workers(), d.Stopper())

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
