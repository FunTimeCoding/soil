package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/store_tester"
	"testing"
	"time"
)

func TestClosedEmptyCandidates(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("empty-closed")
	s.EnsureSession("empty-open")
	s.EnsureSession("logged-closed")
	s.Store.UpdateFields("logged-closed", map[string]any{"lines": 12})
	assert.FatalOnError(t, s.Store.MarkClosed("empty-closed", "logout"))
	assert.FatalOnError(t, s.Store.MarkClosed("logged-closed", "logout"))
	result, e := s.Store.ClosedEmptyCandidates()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "empty-closed", result[0].Identifier)
}

func TestStaleEmptyCandidates(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("stale")
	s.EnsureSession("already-closed")
	s.EnsureSession("spoken")
	s.Store.UpdateFields("spoken", map[string]any{"turn_count": 3})
	assert.FatalOnError(t, s.Store.MarkClosed("already-closed", "logout"))
	now := s.Clock()()
	result, e := s.Store.StaleEmptyCandidates(now.Add(time.Hour))
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "stale", result[0].Identifier)
	none, f := s.Store.StaleEmptyCandidates(now.Add(-time.Hour))
	assert.FatalOnError(t, f)
	assert.Count(t, 0, none)
}

func TestStaleEmptyCandidatesRespectLogActivity(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("touched")
	now := s.Clock()()
	s.Store.UpdateFields(
		"touched",
		map[string]any{"last_active_at": now.Add(time.Hour)},
	)
	result, e := s.Store.StaleEmptyCandidates(now.Add(time.Minute))
	assert.FatalOnError(t, e)
	assert.Count(t, 0, result)
}
