package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

func TestUnkeyedRowStillReachesItsCallsignHolder(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("holder", "Jade", "Jade", "2026-09-01 10:00:00+00:00")
	s.AddQueueRow(
		"Jade",
		"written before the backfill",
		"2026-09-02 10:00:00+00:00",
	)
	drained, e := s.Store.DrainQueue("holder", "Jade")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(t, "written before the backfill", drained[0].Body)
}

func TestUnkeyedRowIsNotReachedWithoutACallsign(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("holder", "Jade", nil, "2026-09-01 10:00:00+00:00")
	s.AddQueueRow(
		"Jade",
		"written before the backfill",
		"2026-09-02 10:00:00+00:00",
	)
	drained, e := s.Store.DrainQueue("holder", "")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}

func TestKeyedRowIsNeverReachedByTheFallback(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("first", "Frost", nil, "2026-09-01 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.PushQueue(
			"first",
			"Frost",
			constant.QueueMessage,
			"for the first Frost",
		),
	)
	s.AddSession("second", "Frost", "Frost", "2026-09-08 10:00:00+00:00")
	drained, e := s.Store.DrainQueue("second", "Frost")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}

func TestUnkeyedPendingEntryIsRetractedByCallsign(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("holder", "Jade", "Jade", "2026-09-01 10:00:00+00:00")
	s.AddQueueRow("Jade", "a stale prompt", "2026-09-02 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.DeletePendingQueue("holder", "Jade", constant.QueueMessage),
	)
	drained, e := s.Store.DrainQueue("holder", "Jade")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
}
