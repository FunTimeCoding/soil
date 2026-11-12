package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestSweepDropsPendingEntryOfSessionClosedPastWeek(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.EditAlias("session-1", "alfa-night")
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"session-1",
			r.Callsign,
			constant.QueueMessage,
			"stale",
		),
	)
	assert.FatalOnError(t, s.Store.Store.MarkClosed("session-1", "exit"))
	s.Store.Advance(8 * 24 * time.Hour)
	s.Roster.RunTimeoutSweep()
	assert.Integer(t, 0, s.Store.PendingQueue("session-1", "stale"))
}

func TestSweepKeepsPendingEntryOfSessionClosedWithinWeek(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.EditAlias("session-1", "alfa-night")
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"session-1",
			r.Callsign,
			constant.QueueMessage,
			"recent",
		),
	)
	assert.FatalOnError(t, s.Store.Store.MarkClosed("session-1", "exit"))
	s.Store.Advance(6 * 24 * time.Hour)
	s.Roster.RunTimeoutSweep()
	assert.Integer(t, 1, s.Store.PendingQueue("session-1", "recent"))
}

func TestSweepDropsPendingEntryOfSessionReleasedAndUnseenPastWeek(t *testing.T) {
	s := service_tester.New(t)
	r := s.Check("session-1")
	s.Store.EditAlias("session-1", "alfa-night")
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"session-1",
			r.Callsign,
			constant.QueueMessage,
			"stale",
		),
	)
	s.Store.ReleaseCallsign(r.Callsign)
	s.Store.Advance(8 * 24 * time.Hour)
	s.Roster.RunTimeoutSweep()
	assert.Integer(t, 0, s.Store.PendingQueue("session-1", "stale"))
}

func TestSweepDropsPendingEntryOfDeletedSession(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Store.PushQueue(
			"session-gone",
			"Ghost",
			constant.QueueMessage,
			"orphan",
		),
	)
	s.Roster.RunTimeoutSweep()
	assert.Integer(t, 0, s.Store.PendingQueue("session-gone", "orphan"))
}

func TestSweepKeepsPendingEntryCarryingNoSession(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddQueueRow("Jade", "unkeyed", "2026-09-02 10:00:00+00:00")
	s.Store.Advance(8 * 24 * time.Hour)
	s.Roster.RunTimeoutSweep()
	assert.Integer(t, 1, s.Store.PendingQueue("", "unkeyed"))
}
