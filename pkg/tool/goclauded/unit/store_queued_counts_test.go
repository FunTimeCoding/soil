package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

func TestQueuedCountsCountWhatWaitsForEachSession(t *testing.T) {
	s := store_tester.New(t)
	first := s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"session-1",
			first.Callsign,
			constant.QueueMessage,
			"one",
		),
	)
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"session-1",
			first.Callsign,
			constant.QueueMessage,
			"two",
		),
	)
	result, e := s.Store.QueuedCounts(s.ListSessions())
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(2), result["session-1"])
	_, found := result["session-2"]
	assert.False(t, found)
}

func TestQueuedCountsCountWhatWasSentBeforeBinding(t *testing.T) {
	s := store_tester.New(t)
	first := s.EnsureSession("session-1")
	s.AddQueueRow(first.Callsign, "sent to the callsign", "2026-11-08 09:00:00")
	result, e := s.Store.QueuedCounts(s.ListSessions())
	assert.FatalOnError(t, e)
	assert.Integer(t, int64(1), result["session-1"])
}

func TestQueuedCountsLeaveOutWhatWasConsumed(t *testing.T) {
	s := store_tester.New(t)
	first := s.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"session-1",
			first.Callsign,
			constant.QueueMessage,
			"read",
		),
	)
	_, e := s.Store.DrainQueue("session-1", first.Callsign)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"session-1",
			first.Callsign,
			constant.QueueMessage,
			"waiting",
		),
	)
	result, f := s.Store.QueuedCounts(s.ListSessions())
	assert.FatalOnError(t, f)
	assert.Integer(t, int64(1), result["session-1"])
}

func TestQueuedCountsWithoutSessions(t *testing.T) {
	s := store_tester.New(t)
	result, e := s.Store.QueuedCounts(nil)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}
