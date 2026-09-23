package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestHistoryCountEmpty(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "bind identity")
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "1 events", result)
}

func TestHistoryCountMultiple(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "one")
	a.Announce(a.Name(), "two")
	a.MustCallTool(
		constant.Update,
		map[string]any{constant.Message: "completed", constant.Topic: "three"},
	)
	result := a.MustCallTool(constant.HistoryCount, map[string]any{})
	assert.StringContains(t, "3 events", result)
}
