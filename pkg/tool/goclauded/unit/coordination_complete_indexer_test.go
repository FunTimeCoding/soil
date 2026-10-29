package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"testing"
)

func TestCompletePushesIndexer(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{constant.Alias: "test-completion"},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "built the thing"},
	)
	pushed := s.Pushed(constant.CompletionCollection)
	assert.Count(t, 1, pushed)
	assert.String(t, "test-completion/1", pushed[0].Path)
	assert.String(t, "built the thing", pushed[0].Body)
	assert.String(t, "session-completion", *pushed[0].SourceType)
}

func TestCompleteWithoutSlugSkipsPush(t *testing.T) {
	s := base.New(t)
	a := s.NewSession(t)
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Complete,
		map[string]any{constant.Message: "completed without slug"},
	)
	assert.Count(t, 0, s.Pushed(constant.CompletionCollection))
}
