package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func consumedAt(
	t *testing.T,
	s *store_tester.Tester,
	body string,
) string {
	t.Helper()
	var result string
	assert.FatalOnError(
		t,
		s.Store.Database().Raw(
			`SELECT COALESCE(CAST(consumed_at AS TEXT), '') FROM queue
			WHERE body = ?`,
			body,
		).Scan(&result).Error,
	)

	return result
}

func TestDrainStampsWhenTheEntryWasDelivered(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("holder", "Wren", constant.QueueMessage, "delivered"),
	)
	assert.String(t, "", consumedAt(t, s, "delivered"))
	drained, e := s.Store.DrainQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, drained)
	assert.String(
		t,
		s.Clock()().Format("2006-01-02 15:04:05.999999999-07:00"),
		consumedAt(t, s, "delivered"),
	)
}

func TestRowConsumedBeforeTheStampExistedIsNotRedelivered(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			`INSERT INTO queue
			(session_identifier, callsign, kind, body, consumed, created_at)
			VALUES ('holder', 'Wren', 'message', 'old', 1,
			'2026-09-02 10:00:00+00:00')`,
		).Error,
	)
	s.Store.StampDelivered()
	drained, e := s.Store.DrainQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.Count(t, 0, drained)
	assert.String(t, "2026-09-02 10:00:00+00:00", consumedAt(t, s, "old"))
}

func TestPeekLeavesTheDeliveryStampUnset(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("holder", "Wren", constant.QueueMessage, "pending"),
	)
	peeked, e := s.Store.PeekQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, peeked)
	assert.String(t, "", consumedAt(t, s, "pending"))
}

func TestDeliveryStampIsUniversal(t *testing.T) {
	s := store_tester.New(t)
	assert.FatalOnError(
		t,
		s.Store.PushQueue("holder", "Wren", constant.QueueMessage, "delivered"),
	)
	_, e := s.Store.DrainQueue("holder", "Wren")
	assert.FatalOnError(t, e)
	assert.True(t, s.Store.UniversalTimestamps())
	assert.StringContains(t, "+00:00", consumedAt(t, s, "delivered"))
}
