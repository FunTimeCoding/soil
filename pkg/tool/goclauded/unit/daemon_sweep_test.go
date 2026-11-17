package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func TestSessionPastItsWindowIsNotSweptWhenDaemonStarts(t *testing.T) {
	o := peer_tester.New(t).DaemonOption()
	s := store_tester.NewOnFile(t, o.LitePath)
	r := s.EnsureSession("session-1")
	s.Announce(r.Callsign, "working through an outage", "")
	later := time.Now().UTC().Add(2 * time.Hour)
	o.Clock = func() time.Time { return later }
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(d.Close)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	assert.String(t, "", s.GetSession("session-1").TimedOut)
}
