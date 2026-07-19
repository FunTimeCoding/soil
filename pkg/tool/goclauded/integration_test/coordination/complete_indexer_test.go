package coordination

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestCompletePushesIndexer(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Slug: "test-completion",
		},
	)
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "built the thing",
		},
	)
	assert.Count(t, 1, s.CompletionIndexer.Pushed)
	assert.String(t, "test-completion/1", s.CompletionIndexer.Pushed[0].Name)
	assert.String(t, "built the thing", s.CompletionIndexer.Pushed[0].Body)
}

func TestCompleteWithoutSlugSkipsPush(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "building things")
	a.MustCallTool(
		constant.Complete,
		map[string]any{
			constant.Message: "completed without slug",
		},
	)
	assert.Count(t, 0, s.CompletionIndexer.Pushed)
}
