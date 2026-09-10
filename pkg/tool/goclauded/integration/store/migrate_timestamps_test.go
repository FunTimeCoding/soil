package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func storedValue(
	t *testing.T,
	s *store_tester.Tester,
	identifier string,
) string {
	t.Helper()
	var result string
	assert.FatalOnError(
		t,
		s.Store.Database().Raw(
			"SELECT CAST(last_seen AS TEXT) FROM session WHERE identifier = ?",
			identifier,
		).Scan(&result).Error,
	)

	return result
}

func setStoredValue(
	t *testing.T,
	s *store_tester.Tester,
	identifier string,
	value string,
) {
	t.Helper()
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			"UPDATE session SET last_seen = ? WHERE identifier = ?",
			value,
			identifier,
		).Error,
	)
}

func TestMigrateTimestampsConvertsAndKeepsPrecision(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("legacy")
	setStoredValue(t, s, "legacy", "2026-08-25 16:28:26.163894+02:00")
	assert.False(t, s.Store.UniversalTimestamps())
	s.Store.NormalizeTimestamps()
	assert.String(
		t,
		"2026-08-25 14:28:26.163894+00:00",
		storedValue(t, s, "legacy"),
	)
	assert.True(t, s.Store.UniversalTimestamps())
}

func TestMigrateTimestampsLeavesUniversalValuesAlone(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("already")
	setStoredValue(t, s, "already", "2026-08-25 14:28:26.163894+00:00")
	s.Store.NormalizeTimestamps()
	assert.String(
		t,
		"2026-08-25 14:28:26.163894+00:00",
		storedValue(t, s, "already"),
	)
}

func TestMigrateTimestampsIsIdempotent(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("twice")
	setStoredValue(t, s, "twice", "2026-08-25 16:28:26.163894+02:00")
	s.Store.NormalizeTimestamps()
	first := storedValue(t, s, "twice")
	s.Store.NormalizeTimestamps()
	assert.String(t, first, storedValue(t, s, "twice"))
}

func TestMigrateTimestampsSkipsUnparseableAndFailsPrecondition(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("broken")
	setStoredValue(t, s, "broken", "not a timestamp")
	s.Store.NormalizeTimestamps()
	assert.String(t, "not a timestamp", storedValue(t, s, "broken"))
	assert.False(t, s.Store.UniversalTimestamps())
}

func TestMigrateTimestampsHandlesTheRepeatedDaylightHour(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("earlier")
	s.EnsureSession("later")
	setStoredValue(t, s, "earlier", "2026-10-25 02:30:00+02:00")
	setStoredValue(t, s, "later", "2026-10-25 02:30:00+01:00")
	s.Store.NormalizeTimestamps()
	assert.String(t, "2026-10-25 00:30:00+00:00", storedValue(t, s, "earlier"))
	assert.String(t, "2026-10-25 01:30:00+00:00", storedValue(t, s, "later"))
	assert.True(t, storedValue(t, s, "earlier") < storedValue(t, s, "later"))
}
