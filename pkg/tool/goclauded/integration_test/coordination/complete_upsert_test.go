package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestCompleteAmend(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "hasty summary",
		},
	)
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "search index and refresh implemented with full test coverage",
		},
	)
	var found int

	for _, c := range s.Store.RecentCompletions() {
		if c.Topic == "search index" && c.Kind == constant.Complete {
			found++
			assert.String(
				t,
				"search index and refresh implemented with full test coverage",
				c.Summary,
			)
		}
	}

	assert.Integer(t, 1, found)
}

func TestCompleteAmendOneEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "first pass",
		},
	)
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "revised with edge cases",
		},
	)
	var found int

	for _, e := range s.Store.Events(event_query.New().Kind(constant.Complete).SetLimit(10)) {
		if e.Metadata[constant.Topic] == "search index" {
			found++
			assert.String(
				t,
				"revised with edge cases",
				e.Metadata[constant.Message],
			)
		}
	}

	assert.Integer(t, 1, found)
}

func TestCompleteDifferentTopics(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done with first",
		},
	)
	a.Announce(a.Name(), "second task")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done with second",
		},
	)
	var count int

	for _, c := range s.Store.RecentCompletions() {
		if c.Kind == constant.Complete {
			count++
		}
	}

	assert.Integer(t, 2, count)
}
