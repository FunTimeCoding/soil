package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestEventSessionID(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Moment,
		map[string]any{
			constant.Line: "a moment",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)

	for _, e := range s.Store.Events(event_query.New().SetLimit(10)) {
		if e.Kind == "register" {
			continue
		}

		assert.True(t, e.SessionIdentifier != "")
	}
}

func TestSummarySessionID(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Summarize,
		map[string]any{
			constant.Body: "session summary",
		},
	)
	v := s.Store.ListSummaries()
	assert.Count(t, 1, v)
	assert.True(t, v[0].SessionIdentifier != "")
}

func TestCompletionSessionID(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	var found bool

	for _, c := range s.Store.RecentCompletions() {
		if c.Kind == constant.Complete {
			assert.True(t, c.SessionIdentifier != "")
			found = true
		}
	}

	assert.True(t, found)
}
