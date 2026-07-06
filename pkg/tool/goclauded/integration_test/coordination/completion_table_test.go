package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestCompletionTableAfterComplete(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "search index")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "search index implemented",
		},
	)
	c := s.Store.RecentCompletions()
	assert.Count(t, 1, c)
	assert.String(t, a.Name(), c[0].Name)
	assert.String(t, constant.Complete, c[0].Kind)
	assert.String(t, "search index", c[0].Topic)
	assert.String(t, "search index implemented", c[0].Summary)
}

func TestCompletionTableAfterUpdate(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "initial")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone reached",
		},
	)
	c := s.Store.RecentCompletions()
	assert.Count(t, 1, c)
	assert.String(t, a.Name(), c[0].Name)
	assert.String(t, constant.Update, c[0].Kind)
	assert.String(t, "milestone reached", c[0].Topic)
	assert.String(t, "completed", c[0].Summary)
}

func TestCompletionTableHookContext(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "search index")
	b.CheckLive()
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	a.Announce(a.Name(), "next task")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "pivot to testing",
		},
	)
	r := b.CheckLive()
	assert.True(t, r.Changed)
	completions := clientEntriesByKind(
		r.Entries,
		constant.QueueSessionComplete,
	)
	updates := clientEntriesByKind(r.Entries, constant.QueueSessionUpdate)
	assert.True(t, len(completions) > 0)
	assert.True(t, len(updates) > 0)
}
