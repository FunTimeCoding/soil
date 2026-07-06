package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestHistoryFormatsIDs(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "test topic")
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "announced: test topic", history)
}

func TestHistoryLimit(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "topic-alpha")
	a.Announce(a.Name(), "topic-bravo")
	a.Announce(a.Name(), "topic-charlie")
	history := a.MustCallTool(
		constant.History,
		map[string]any{constant.Limit: float64(1)},
	)
	assert.StringContains(t, "topic-charlie", history)
	assert.StringNotContains(t, "topic-alpha", history)
}

func TestHistorySkip(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "topic-alpha")
	a.Announce(a.Name(), "topic-bravo")
	a.Announce(a.Name(), "topic-charlie")
	history := a.MustCallTool(
		constant.History,
		map[string]any{
			constant.Limit:  float64(1),
			constant.Offset: float64(1),
		},
	)
	assert.StringContains(t, "topic-bravo", history)
	assert.StringNotContains(t, "topic-charlie", history)
}

func TestHistoryAllKinds(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.Update,
		map[string]any{
			constant.Message: "completed",
			constant.Topic:   "milestone",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "finished",
		},
	)
	history := a.MustCallTool(constant.History, map[string]any{})
	assert.StringContains(t, "announced: working", history)
	assert.StringContains(t, "updated milestone: completed", history)
	assert.StringContains(t, "completed milestone: finished", history)
}
