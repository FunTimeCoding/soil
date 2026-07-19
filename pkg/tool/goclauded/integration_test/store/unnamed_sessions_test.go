package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestUnnamedSessionsFilterAndOrder(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	s.EnsureSession("session-3")
	s.Store.UpdateFields(
		"session-1",
		map[string]any{"session_timestamp": "2026-07-03T10:00:00Z"},
	)
	s.Store.UpdateFields(
		"session-2",
		map[string]any{"session_timestamp": "2026-07-01T10:00:00Z"},
	)
	s.Store.UpdateFields(
		"session-3",
		map[string]any{"session_timestamp": "2026-07-02T10:00:00Z"},
	)
	editAlias(s, "session-3", "charted")
	sessions, e := s.Store.UnnamedSessions(0, 0)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, sessions)
	assert.String(t, "session-2", sessions[0].Identifier)
	assert.String(t, "session-1", sessions[1].Identifier)
}

func TestUnnamedSessionsLimit(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	s.Store.UpdateFields(
		"session-1",
		map[string]any{"session_timestamp": "2026-07-01T10:00:00Z"},
	)
	s.Store.UpdateFields(
		"session-2",
		map[string]any{"session_timestamp": "2026-07-02T10:00:00Z"},
	)
	sessions, e := s.Store.UnnamedSessions(1, 1)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, sessions)
	assert.String(t, "session-2", sessions[0].Identifier)
}
