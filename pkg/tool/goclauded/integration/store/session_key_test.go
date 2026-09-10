package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func addSession(
	t *testing.T,
	s *store_tester.Tester,
	identifier string,
	name string,
	callsign any,
	started string,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			`INSERT INTO session (identifier, name, callsign, started_at)
			VALUES (?, ?, ?, ?)`,
			identifier,
			name,
			callsign,
			started,
		).Error,
	)
}

func addQueueRow(
	t *testing.T,
	s *store_tester.Tester,
	callsign any,
	body string,
	created string,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			`INSERT INTO queue (callsign, kind, body, consumed, created_at)
			VALUES (?, 'message', ?, 0, ?)`,
			callsign,
			body,
			created,
		).Error,
	)
}

func sessionKeyFor(
	t *testing.T,
	s *store_tester.Tester,
	body string,
) string {
	t.Helper()
	var result string
	assert.FatalOnError(
		t,
		s.Store.Database().Raw(
			`SELECT COALESCE(session_identifier, '') FROM queue
			WHERE body = ?`,
			body,
		).Scan(&result).Error,
	)

	return result
}

func queueRows(
	t *testing.T,
	s *store_tester.Tester,
) int64 {
	t.Helper()
	var result int64
	assert.FatalOnError(
		t,
		s.Store.Database().Raw("SELECT COUNT(*) FROM queue").Scan(
			&result,
		).Error,
	)

	return result
}

func TestSessionKeyAttributesToTheMostRecentPriorHolder(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "early", "Ash", nil, "2026-09-01 10:00:00+00:00")
	addSession(t, s, "late", "Ash", "Ash", "2026-09-03 10:00:00+00:00")
	addQueueRow(t, s, "Ash", "before both", "2026-08-30 10:00:00+00:00")
	addQueueRow(t, s, "Ash", "between them", "2026-09-02 10:00:00+00:00")
	addQueueRow(t, s, "Ash", "after both", "2026-09-04 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.String(t, "early", sessionKeyFor(t, s, "between them"))
	assert.String(t, "late", sessionKeyFor(t, s, "after both"))
	assert.String(t, "", sessionKeyFor(t, s, "before both"))
	assert.Integer(t, 2, queueRows(t, s))
}

func TestSessionKeyAttributesRowsOfAReleasedSession(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "released", "Frost", nil, "2026-09-01 10:00:00+00:00")
	addQueueRow(t, s, "Frost", "for a freed name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.String(t, "released", sessionKeyFor(t, s, "for a freed name"))
	assert.Integer(t, 1, queueRows(t, s))
}

func TestSessionKeyDropsRowsItCannotAttribute(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "nameless", "", nil, "2026-09-01 10:00:00+00:00")
	addQueueRow(t, s, "", "empty callsign", "2026-09-02 10:00:00+00:00")
	addQueueRow(t, s, nil, "no callsign", "2026-09-02 10:00:00+00:00")
	addQueueRow(t, s, "Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	assert.Integer(t, 0, queueRows(t, s))
}

func TestSessionKeyIsIdempotent(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	addQueueRow(t, s, "Wren", "a message", "2026-09-02 10:00:00+00:00")
	addQueueRow(t, s, "Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	s.Store.BackfillSessionKey()
	first := sessionKeyFor(t, s, "a message")
	rows := queueRows(t, s)
	s.Store.BackfillSessionKey()
	assert.String(t, first, sessionKeyFor(t, s, "a message"))
	assert.Integer(t, rows, queueRows(t, s))
	assert.Integer(t, 1, queueRows(t, s))
}

func TestSessionKeyDeclinesOnUnnormalizedTimestamps(t *testing.T) {
	s := store_tester.New(t)
	addSession(t, s, "holder", "Wren", "Wren", "2026-09-01 10:00:00+00:00")
	addQueueRow(t, s, "Wren", "a message", "2026-09-02 10:00:00+00:00")
	addQueueRow(t, s, "Nobody", "unheld name", "2026-09-02 10:00:00+00:00")
	setStoredValue(t, s, "holder", "2026-09-01 12:00:00+02:00")
	assert.False(t, s.Store.UniversalTimestamps())
	s.Store.BackfillSessionKey()
	assert.String(t, "", sessionKeyFor(t, s, "a message"))
	assert.Integer(t, 2, queueRows(t, s))
}
