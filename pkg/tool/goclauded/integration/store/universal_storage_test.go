package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestEveryWriterStoresUniversalTime(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.Announce(r.Callsign, "a topic", ""))
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Announce,
			r.Callsign,
			map[string]string{constant.Topic: "a topic"},
		),
	)
	_, e := s.Store.UpsertCompletion(
		"session-1",
		r.Callsign,
		constant.Complete,
		"a topic",
		"a summary",
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.Store.UpsertSummary("session-1", r.Callsign, "a summary"),
	)
	assert.FatalOnError(
		t,
		s.Store.SendPulse("session-1", r.Callsign, "a pulse"),
	)
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"session-1",
			r.Callsign,
			constant.QueueTimeout,
			"an entry",
		),
	)
	assert.FatalOnError(
		t,
		s.Store.SendNotification(
			"session-1",
			r.Callsign,
			"source",
			"a notification",
		),
	)
	assert.FatalOnError(
		t,
		s.Store.SendMessage("Someone", r.Callsign, "a message"),
	)
	assert.FatalOnError(t, s.Store.MarkClosed("session-1", "logout"))
	assert.True(t, s.Store.UniversalTimestamps())
}
