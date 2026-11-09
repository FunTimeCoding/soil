package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

func TestSummarySessionMigrationKeepsNewestRow(t *testing.T) {
	s := store_tester.NewOnFile(t, seedDuplicateSummaries(t))
	summaries := s.ListSummaries()
	assert.Count(t, 1, summaries)
	assert.String(t, "rewrite", summaries[0].Body)
}

func TestSummarySessionRefusesSecondRow(t *testing.T) {
	s := store_tester.New(t)
	_, e := s.Store.EnsureSession("session-1")
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.Store.UpsertSummary("session-1", "Frost", "rewrite"),
	)
	assert.Error(
		t,
		s.Mapper().Exec(
			"INSERT INTO summary (session_identifier, name, body) VALUES (?, ?, ?)",
			"session-1",
			"Frost",
			"first draft",
		).Error,
	)
}
