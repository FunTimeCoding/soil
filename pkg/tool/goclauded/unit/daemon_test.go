package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/credential"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/worker"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestDaemonClosesBothPeerClientsOnce(t *testing.T) {
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

func TestUtilizationFallbackGuardsItsPeer(t *testing.T) {
	o, _, _ := newDaemonOption(t)
	u := mock_client.New()
	o.UtilizationFallback = true
	o.Utilization = func() face.AnthropicClient { return u }
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	workers := d.Workers()
	assert.Count(t, 13, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 3, guard)
	}

	d.Close()
	d.Close()
	assert.Integer(t, 1, u.Closed)
}

func TestUtilizationClockReadsOptionsCredential(t *testing.T) {
	o, _, _ := newDaemonOption(t)
	u := mock_client.New()
	u.Result = utilization.Parse(
		[]byte(fixture.Read("claude", "utilization.json")),
	)
	o.UtilizationFallback = true
	o.Utilization = func() face.AnthropicClient { return u }
	o.Credential = func() *credential.Credential {
		return credential.New("alfa-token", "", "", time.Now().Add(time.Hour))
	}
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)

	for _, w := range d.Workers() {
		if c, okay := w.(*worker.Worker); okay {
			c.Poll()
		}
	}

	assert.Integer(t, 1, u.Reads)
	r, e := store_tester.NewOnFile(t, o.LitePath).Store.LatestRateSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, r)
	assert.Integer(t, 22, r.SevenDayPercent)
}

func TestDaemonHandsOutItsGuardsFirst(t *testing.T) {
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

func TestDaemonClearsStaleBindingsBeforeItServes(t *testing.T) {
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

func TestDaemonLeavesNothingRunning(t *testing.T) {
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
