package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestReconcileCountsPushedSummary(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("session-1", "Ash", "indexed the corpus"),
	)
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.Float(
		t,
		1,
		sample(
			t,
			s.Registry,
			"goclauded_reconciled_total",
			map[string]string{
				constant.Collection: constant.SummaryCollection,
				constant.Outcome:    constant.OutcomePushed,
			},
		),
	)
	assert.Float(
		t,
		1,
		sample(
			t,
			s.Registry,
			"goclauded_indexed_documents",
			map[string]string{constant.Collection: constant.SummaryCollection},
		),
	)
}

func TestReconcileRepushesSummaryWhoseMetadataChanged(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("session-1", "Ash", "indexed the corpus"),
	)
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.FatalOnError(
		t,
		s.Store.Store.Database().Exec(
			`UPDATE session SET started_at = '2026-09-02 10:00:00+00:00'
			WHERE identifier = 'session-1'`,
		).Error,
	)
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.Count(t, 2, s.Query.Pushed)
}

func TestReconcileLeavesCurrentSummaryAlone(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("session-1", "Ash", "indexed the corpus"),
	)
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.Count(t, 0, s.Document.ReconcileSummaries())
	assert.Count(t, 1, s.Query.Pushed)
	assert.Float(
		t,
		1,
		sample(
			t,
			s.Registry,
			"goclauded_reconciled_total",
			map[string]string{
				constant.Collection: constant.SummaryCollection,
				constant.Outcome:    constant.OutcomePushed,
			},
		),
	)
	assert.Float(
		t,
		1,
		sample(
			t,
			s.Registry,
			"goclauded_indexed_documents",
			map[string]string{constant.Collection: constant.SummaryCollection},
		),
	)
}
