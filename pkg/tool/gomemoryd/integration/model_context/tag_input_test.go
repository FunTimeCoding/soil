package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
	"testing"
)

func TestTagInputStripsNotationArray(t *testing.T) {
	s := model_context_tester.NewTagged(t)
	result := s.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.ReplaceAll:       `["build","go-conventions"]`,
		},
	)
	assert.StringContains(t, "Stripped", result)
	assert.StringContains(t, "saved as: build, go-conventions", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, m.Tags)
	assert.String(t, "build", m.Tags[0])
}

func TestTagInputStaysQuietWhenClean(t *testing.T) {
	s := model_context_tester.NewTagged(t)
	result := s.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.ReplaceAll:       "build,go-conventions",
		},
	)
	assert.StringNotContains(t, "Stripped", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, m.Tags)
}
