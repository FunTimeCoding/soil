//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/model_context_tester"
	"testing"
)

func TestCreateMemory(t *testing.T) {
	s := model_context_tester.New(t)
	result := s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "test content",
			constant.Description: "a test",
		},
	)
	assert.StringContains(t, "Created memory", result)
	memories, e := s.Store().ListMemories("", "", true)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, memories)
}

func TestUpdateMemory(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "original",
			constant.Description: "a test",
		},
	)
	result := s.MustCallTool(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.MemoryName:       "test memory",
			constant.Content:          "updated",
			constant.Description:      "a test",
		},
	)
	assert.StringContains(t, "Updated memory", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "updated", m.Content)
}

func TestUpdateMemoryRequiresID(t *testing.T) {
	s := model_context_tester.New(t)
	result := s.MustCallToolError(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryName:  "test",
			constant.Content:     "test",
			constant.Description: "test",
		},
	)
	assert.StringContains(t, "input schema validation failed", result)
	assert.StringContains(t, "Missing:[memory_id]", result)
}

func TestUpdateMemoryWithWrongParameterName(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "original",
			constant.Description: "a test",
		},
	)
	result := s.MustCallToolError(
		constant.UpdateMemory,
		map[string]any{
			"id":                 1,
			constant.MemoryName:  "test memory",
			constant.Content:     "should fail",
			constant.Description: "a test",
		},
	)
	assert.StringContains(t, "input schema validation failed", result)
	assert.StringContains(t, "Missing:[memory_id]", result)
	assert.StringContains(t, "Properties:[id]", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "original", m.Content)
}
