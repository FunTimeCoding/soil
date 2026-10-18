package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
	"time"
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
			"relay",
			"a notification",
		),
	)
	s.SendMessage("Someone", r.Callsign, "a message")
	writeCruise(t, s.Store, "session-1", r.Callsign)
	assert.FatalOnError(t, s.Store.MarkClosed("session-1", "logout"))
	assert.True(t, s.Store.UniversalTimestamps())
}

func TestCruiseWritersStoreUniversalTimeFromALocalClock(t *testing.T) {
	berlin, e := time.LoadLocation("Europe/Berlin")
	assert.FatalOnError(t, e)
	s := store.New(
		lite.NewMemory(),
		func() time.Time { return time.Now().In(berlin) },
	)
	defer s.Close()
	writeCruise(t, s, "session-1", "Alfa")
	assert.True(t, s.UniversalTimestamps())
}
