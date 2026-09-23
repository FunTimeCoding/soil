package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestSessionKeyAttributesToTheMostRecentPriorHolder(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("early", "Ash", nil, "2026-09-01 10:00:00+00:00")
	s.AddSession("late", "Ash", "Ash", "2026-09-03 10:00:00+00:00")
	s.AddQueueRow("Ash", "before both", "2026-08-30 10:00:00+00:00")
	s.AddQueueRow("Ash", "between them", "2026-09-02 10:00:00+00:00")
	s.AddQueueRow("Ash", "after both", "2026-09-04 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.String(t, "early", s.SessionKeyFor("between them"))
	assert.String(t, "late", s.SessionKeyFor("after both"))
	assert.String(t, "", s.SessionKeyFor("before both"))
	assert.Integer(t, 2, s.QueueRows())
}

func TestSessionKeyAttributesRowsOfAReleasedSession(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("released", "Frost", nil, "2026-09-01 10:00:00+00:00")
	s.AddQueueRow("Frost", "for a freed name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.String(t, "released", s.SessionKeyFor("for a freed name"))
	assert.Integer(t, 1, s.QueueRows())
}

func TestSessionKeyDropsRowsItCannotAttribute(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("nameless", "", nil, "2026-09-01 10:00:00+00:00")
	s.AddQueueRow("", "empty callsign", "2026-09-02 10:00:00+00:00")
	s.AddQueueRow(nil, "no callsign", "2026-09-02 10:00:00+00:00")
	s.AddQueueRow("Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.Integer(t, 0, s.QueueRows())
}

func TestSessionKeyIsIdempotent(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	s.AddQueueRow("Wren", "a message", "2026-09-02 10:00:00+00:00")
	s.AddQueueRow("Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	first := s.SessionKeyFor("a message")
	rows := s.QueueRows()
	s.Store.BackfillSessionKey()
	assert.String(t, first, s.SessionKeyFor("a message"))
	assert.Integer(t, rows, s.QueueRows())
	assert.Integer(t, 1, s.QueueRows())
}

func TestSessionKeyDeclinesOnUnnormalizedTimestamps(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	s.AddQueueRow("Wren", "a message", "2026-09-02 10:00:00+00:00")
	s.AddQueueRow("Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	s.SetStoredValue("holder", "2026-09-01 12:00:00+02:00")
	assert.False(t, s.Store.UniversalTimestamps())
	s.Store.BackfillSessionKey()
	assert.String(t, "", s.SessionKeyFor("a message"))
	assert.Integer(t, 2, s.QueueRows())
}
