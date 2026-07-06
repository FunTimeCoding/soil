package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestComplete(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building search index")
	a.CheckLive()
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "search index and refresh implemented",
		},
	)
	r := a.CheckLive()
	completions := clientEntriesByKind(
		r.Entries,
		constant.QueueSessionComplete,
	)
	assert.True(t, len(completions) > 0)
	assert.StringContains(t, "building search index", completions[0].Body)
}

func TestCompleteHistoryEvent(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "reviewing proposals")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "reviewed 8 proposals into design docs",
		},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "completed reviewing proposals", history)
	assert.StringContains(t, "reviewed 8 proposals into design docs", history)
}

func TestCompleteAndReannounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "first task")
	a.CheckLive()
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done with first",
		},
	)
	a.Announce(a.Name(), "second task")
	r := a.CheckLive()
	announces := clientEntriesByKind(r.Entries, constant.QueueSessionAnnounce)
	assert.True(t, len(announces) > 0)
	assert.StringContains(t, "second task", announces[len(announces)-1].Body)
}

func TestCompleteWithExplicitTopic(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "main work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "main work done",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Topic:   "ad-hoc fix",
			constant.Message: "fixed the spacing bug",
		},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "completed ad-hoc fix", history)
	assert.StringContains(t, "fixed the spacing bug", history)
}

func TestCompleteNoTopicNoAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "done",
		},
	)
	result := a.MustCallToolError(
		constant.Complete,
		map[string]any{
			constant.Message: "nothing to complete",
		},
	)
	assert.StringContains(t, "no active topic", result)
}

func TestCompleteBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	result := a.MustCallToolError(
		constant.Complete,
		map[string]any{
			constant.Message: "nothing to complete",
		},
	)
	assert.StringContains(t, "announce first", result)
}
