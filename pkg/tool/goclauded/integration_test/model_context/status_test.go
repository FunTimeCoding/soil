package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestStatusWithTopic(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building search index")
	result := a.MustCallTool(constant.Status, map[string]any{})
	assert.StringContains(t, a.Name(), result)
	assert.StringContains(t, "building search index", result)
}

func TestStatusAfterComplete(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "some work")
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "done"},
	)
	result := a.MustCallTool(constant.Status, map[string]any{})
	assert.StringContains(t, "(none)", result)
}

func TestStatusBeforeAnnounce(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	result := a.MustCallToolError(constant.Status, map[string]any{})
	assert.StringContains(t, "announce first", result)
}
