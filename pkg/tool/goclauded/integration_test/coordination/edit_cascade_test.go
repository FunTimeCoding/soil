package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestEditCascadeSummary(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "original summary",
		},
	)
	events := s.Store.Events(event_query.New().Kind(constant.Summarize).SetLimit(1))
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.EditEvent,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected summary via edit",
		},
	)
	summaries := s.Store.ListSummaries()
	assert.Count(t, 1, summaries)
	assert.String(t, "corrected summary via edit", summaries[0].Body)
}

func TestEditCascadeSummaryPushesIndexer(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Slug: "test-session",
		},
	)
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "original",
		},
	)
	before := len(s.SummaryIndexer.Pushed)
	events := s.Store.Events(event_query.New().Kind(constant.Summarize).SetLimit(1))
	a.MustCallTool(
		constant.EditEvent,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "edited via cascade",
		},
	)
	assert.Integer(t, before+1, len(s.SummaryIndexer.Pushed))
	last := s.SummaryIndexer.Pushed[len(s.SummaryIndexer.Pushed)-1]
	assert.String(t, "edited via cascade", last.Body)
}

func TestEditCascadeCompletion(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "hasty completion",
		},
	)
	events := s.Store.Events(event_query.New().Kind(constant.Complete).SetLimit(1))
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.EditEvent,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "corrected completion with details",
		},
	)
	completions := s.Store.RecentCompletions()
	var found bool

	for _, c := range completions {
		if c.Topic == "search index" && c.Kind == constant.Complete {
			assert.String(t, "corrected completion with details", c.Summary)
			found = true
		}
	}

	assert.True(t, found)
}

func TestEditMomentNoCascade(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Moment,
		map[string]any{
			constant.Line: "something landed",
		},
	)
	before := len(s.SummaryIndexer.Pushed)
	events := s.Store.Events(event_query.New().Kind(constant.Moment).SetLimit(1))
	assert.Count(t, 1, events)
	a.MustCallTool(
		constant.EditEvent,
		map[string]any{
			constant.Identifier: float64(events[0].Identifier),
			constant.Message:    "revised moment",
		},
	)
	assert.Integer(t, before, len(s.SummaryIndexer.Pushed))
	edited := s.Store.Events(event_query.New().Kind(constant.Moment).SetLimit(1))
	assert.String(t, "revised moment", edited[0].Metadata[constant.Line])
}
