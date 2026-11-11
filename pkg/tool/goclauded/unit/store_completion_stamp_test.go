package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"path/filepath"
	"testing"
	"time"
)

func TestRecompletionMovesCompletionTime(t *testing.T) {
	s := store_tester.New(t)
	s.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	first, e := s.Store.UpsertCompletion(
		"session-1",
		"Ash",
		constant.Complete,
		"search index",
		"first pass",
	)
	assert.FatalOnError(t, e)
	s.Advance(time.Hour)
	second, f := s.Store.UpsertCompletion(
		"session-1",
		"Ash",
		constant.Complete,
		"search index",
		"second pass",
	)
	assert.FatalOnError(t, f)
	assert.Duration(t, time.Hour, second.CompletedAt.Sub(first.CompletedAt))
	stored := s.SessionCompletions("session-1")
	assert.Count(t, 1, stored)
	assert.True(t, stored[0].CompletedAt.Equal(second.CompletedAt))
	assert.True(t, stored[0].CreatedAt.Equal(first.CreatedAt))
}

func TestCompletionWithoutStampTakesCreationTime(t *testing.T) {
	path := filepath.Join(t.TempDir(), "goclauded.db")
	s := store_tester.NewOnFile(t, path)
	s.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	assert.FatalOnError(
		t,
		s.Store.Database().Exec(
			`INSERT INTO completion
			(session_identifier, name, kind, topic, summary, sequence, created_at)
			VALUES ('session-1', 'Ash', 'complete', 'search index', 'old', 1,
			'2026-09-02 10:00:00+00:00')`,
		).Error,
	)
	stored := store_tester.NewOnFile(t, path).SessionCompletions("session-1")
	assert.Count(t, 1, stored)
	assert.True(t, stored[0].CompletedAt.Equal(stored[0].CreatedAt))
	assert.False(t, stored[0].CompletedAt.IsZero())
}
