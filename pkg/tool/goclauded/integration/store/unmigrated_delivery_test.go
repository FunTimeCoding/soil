package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestUnkeyedRowStillReachesItsCallsignHolder(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	addQueueRow(
		t,
		s,
		"Wren",
		"written before the backfill",
		"2026-09-02 10:00:00+00:00",
	)
	drained, e := s.Store.DrainQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(t, "written before the backfill", drained[0].Body)
}

func TestUnkeyedRowIsNotReachedWithoutACallsign(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "holder", "Wren", nil, "2026-09-01 10:00:00+00:00")
	addQueueRow(
		t,
		s,
		"Wren",
		"written before the backfill",
		"2026-09-02 10:00:00+00:00",
	)
	drained, e := s.Store.DrainQueue("holder", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}

func TestKeyedRowIsNeverReachedByTheFallback(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "first", "Frost", nil, "2026-09-01 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"first",
			"Frost",
			constant.QueueMessage,
			"for the first Frost",
		),
	)
	addSession(t, s, "second", "Frost", "Frost", "2026-09-08 10:00:00+00:00")
	drained, e := s.Store.DrainQueue("second", "Frost")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}

func TestUnkeyedPendingEntryIsRetractedByCallsign(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	addQueueRow(t, s, "Wren", "a stale prompt", "2026-09-02 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.DeletePendingQueue("holder", "Wren", constant.QueueMessage),
	)
	drained, e := s.Store.DrainQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}
