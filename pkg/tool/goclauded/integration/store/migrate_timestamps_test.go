package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
)

func TestMigrateTimestampsConvertsAndKeepsPrecision(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("legacy")
	s.SetStoredValue("legacy", "2026-08-25 16:28:26.163894+02:00")
	assert.False(t, s.Store.UniversalTimestamps())
	s.Store.NormalizeTimestamps()
	assert.String(
		t,
		"2026-08-25 14:28:26.163894+00:00",
		s.StoredValue("legacy"),
	)
	assert.True(t, s.Store.UniversalTimestamps())
}

func TestMigrateTimestampsLeavesUniversalValuesAlone(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("already")
	s.SetStoredValue("already", "2026-08-25 14:28:26.163894+00:00")
	s.Store.NormalizeTimestamps()
	assert.String(
		t,
		"2026-08-25 14:28:26.163894+00:00",
		s.StoredValue("already"),
	)
}

func TestMigrateTimestampsIsIdempotent(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("twice")
	s.SetStoredValue("twice", "2026-08-25 16:28:26.163894+02:00")
	s.Store.NormalizeTimestamps()
	first := s.StoredValue("twice")
	s.Store.NormalizeTimestamps()
	assert.String(t, first, s.StoredValue("twice"))
}

func TestMigrateTimestampsSkipsUnparseableAndFailsPrecondition(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("broken")
	s.SetStoredValue("broken", "not a timestamp")
	s.Store.NormalizeTimestamps()
	assert.String(t, "not a timestamp", s.StoredValue("broken"))
	assert.False(t, s.Store.UniversalTimestamps())
}

func TestMigrateTimestampsHandlesTheRepeatedDaylightHour(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("earlier")
	s.EnsureSession("later")
	s.SetStoredValue("earlier", "2026-10-25 02:30:00+02:00")
	s.SetStoredValue("later", "2026-10-25 02:30:00+01:00")
	s.Store.NormalizeTimestamps()
	assert.String(t, "2026-10-25 00:30:00+00:00", s.StoredValue("earlier"))
	assert.String(t, "2026-10-25 01:30:00+00:00", s.StoredValue("later"))
	assert.True(t, s.StoredValue("earlier") < s.StoredValue("later"))
}
