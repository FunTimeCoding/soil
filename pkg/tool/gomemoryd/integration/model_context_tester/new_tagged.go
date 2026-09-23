package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"testing"
)

func NewTagged(t *testing.T) *Tester {
	t.Helper()
	result := New(t)
	result.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alfa",
			constant.Content:     "first",
			constant.Description: "a test",
		},
	)

	return result
}
