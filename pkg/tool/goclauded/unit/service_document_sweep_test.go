package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
)

func TestReconcileSweepsCompletionUnderFormerName(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	s.Complete("session-1", "Ash", "search index", "indexed the corpus")
	s.Store.EditAlias("session-1", "ash-blair")
	assert.Count(t, 0, s.Document.ReconcileCompletions())
	assert.Count(t, 1, s.Query.Deleted)
	assert.String(t, "completions", s.Query.Deleted[0].Collection)
	assert.String(t, "ash-night/1", s.Query.Deleted[0].Path)
	pushed := s.Pushed(constant.CompletionCollection)
	assert.String(t, "ash-blair/1", pushed[len(pushed)-1].Path)
	assert.Float(
		t,
		1,
		sample(
			t,
			s.Registry,
			"goclauded_reconciled_total",
			map[string]string{
				constant.Collection: constant.CompletionCollection,
				constant.Outcome:    constant.OutcomeSwept,
			},
		),
	)
	assert.Count(t, 0, s.Document.ReconcileCompletions())
	assert.Count(t, 1, s.Query.Deleted)
}

func TestReconcileSweepsSummaryUnderFormerName(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("session-1", "Ash", "indexed the corpus"),
	)
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	s.Store.EditAlias("session-1", "ash-blair")
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.Count(t, 1, s.Query.Deleted)
	assert.String(t, "summaries", s.Query.Deleted[0].Collection)
	assert.String(t, "ash-night", s.Query.Deleted[0].Path)
}

func TestReconcileLeavesDocumentOfUnknownSession(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	s.Complete("session-1", "Ash", "search index", "indexed the corpus")
	stranger := map[string][]string{"session_id": {"session-2"}}
	_, e := s.Query.PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "completions",
			Path:       "stranger/1",
			Body:       "not ours to judge",
			Metadata:   &stranger,
		},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, s.Document.ReconcileCompletions())
	assert.Count(t, 0, s.Query.Deleted)
}
