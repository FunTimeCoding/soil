package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"testing"
	"time"
)

func TestRecompletionPushesLaterCompletionTime(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	s.Complete("session-1", "Ash", "search index", "first pass")
	s.Store.Advance(time.Hour)
	s.Complete("session-1", "Ash", "search index", "second pass")
	pushed := s.Pushed(constant.CompletionCollection)
	assert.Count(t, 2, pushed)
	assert.Duration(
		t,
		time.Hour,
		pushedCompletionTime(t, pushed[1]).Sub(
			pushedCompletionTime(t, pushed[0]),
		),
	)
	assert.Strings(
		t,
		(*pushed[0].Metadata)["created_at"],
		(*pushed[1].Metadata)["created_at"],
	)
}

func TestEditEventKeepsCompletionTime(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	s.Complete("session-1", "Ash", "search index", "hasty completion")
	s.Store.Advance(time.Hour)
	events := s.Store.Events(
		event_query.New().Kind(constant.Complete).SetLimit(1),
	)
	_, e := s.Event.EditEvent(events[0].Identifier, "corrected completion")
	assert.FatalOnError(t, e)
	pushed := s.Pushed(constant.CompletionCollection)
	assert.Count(t, 2, pushed)
	assert.String(t, "corrected completion", pushed[1].Body)
	assert.Strings(
		t,
		(*pushed[0].Metadata)[constant.CompletedAtMetadata],
		(*pushed[1].Metadata)[constant.CompletedAtMetadata],
	)
}

func TestReconcileRepushesCompletionMissingTime(t *testing.T) {
	s := service_tester.New(t)
	s.Store.AddSession("session-1", "Ash", "Ash", "2026-09-01 10:00:00+00:00")
	s.Store.EditAlias("session-1", "ash-night")
	s.Complete("session-1", "Ash", "search index", "indexed the corpus")
	stale := map[string][]string{"created_at": {"2026-09-01T10:00:00Z"}}
	_, e := s.Query.PostDocument(
		t.Context(),
		client.PostDocumentJSONRequestBody{
			Collection: "completions",
			Path:       "ash-night/1",
			Body:       "indexed the corpus",
			Metadata:   &stale,
		},
	)
	assert.FatalOnError(t, e)
	before := len(s.Pushed(constant.CompletionCollection))
	assert.Count(t, 0, s.Document.ReconcileCompletions())
	assert.Integer(t, before+1, len(s.Pushed(constant.CompletionCollection)))
	assert.Count(t, 0, s.Document.ReconcileCompletions())
	assert.Integer(t, before+1, len(s.Pushed(constant.CompletionCollection)))
}
