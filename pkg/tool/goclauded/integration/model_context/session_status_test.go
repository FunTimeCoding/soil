package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/base"
	"testing"
)

func TestSessionStatusWithTopic(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building search index")
	result := a.MustCallTool(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, a.Name(), result)
	assert.StringContains(t, "building search index", result)
}

func TestSessionStatusAfterComplete(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "some work")
	a.MustCallTool(constant.Complete, map[string]any{constant.Message: "done"})
	result := a.MustCallTool(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, "(none)", result)
}

func TestSessionStatusBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	result := a.MustCallToolError(constant.SessionStatus, map[string]any{})
	assert.StringContains(t, "announce first", result)
}
