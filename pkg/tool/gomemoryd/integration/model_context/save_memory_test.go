package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
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
	memories, e := s.Store().ListMemories("", "", "", true)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, memories)
}

func TestCreateMemoryWithTags(t *testing.T) {
	s := model_context_tester.New(t)
	result := s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "test content",
			constant.Description: "a test",
			constant.Tags:        "build,groom",
		},
	)
	assert.StringContains(t, "Created memory", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"build", "groom"}, m.Tags)
}

func TestCreateMemoryStripsTagCruft(t *testing.T) {
	s := model_context_tester.New(t)
	result := s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "test content",
			constant.Description: "a test",
			constant.Tags:        `["build", "groom"]`,
		},
	)
	assert.StringContains(t, "Stripped", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"build", "groom"}, m.Tags)
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

func TestCreateMemoryWithBase(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "See `fleet.md`.",
			constant.Description: "a test",
			constant.Base:        "doc/ai/runbook",
		},
	)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "doc/ai/runbook", m.Metadata[constant.BaseKey])
}

func TestUpdateMemoryKeepsBaseWhenOmitted(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "original",
			constant.Description: "a test",
			constant.Base:        "doc/ai/runbook",
		},
	)
	s.MustCallTool(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.MemoryName:       "test memory",
			constant.Content:          "updated",
			constant.Description:      "a test",
		},
	)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "doc/ai/runbook", m.Metadata[constant.BaseKey])
}

func TestUpdateMemoryKeepsBaseWhenEmpty(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "original",
			constant.Description: "a test",
			constant.Base:        "doc/ai/runbook",
		},
	)
	s.MustCallTool(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.Content:          "updated",
			constant.Base:             "",
		},
	)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "doc/ai/runbook", m.Metadata[constant.BaseKey])
}

func TestUpdateMemoryClearsBaseWithClearBase(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "test memory",
			constant.Content:     "original",
			constant.Description: "a test",
			constant.Base:        "doc/ai/runbook",
		},
	)
	s.MustCallTool(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.ClearBase:        true,
		},
	)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	_, found := m.Metadata[constant.BaseKey]
	assert.Boolean(t, false, found)
	assert.String(t, "original", m.Content)
}

func TestUpdateMemoryWithDescriptionOnly(t *testing.T) {
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
			constant.Description:      "a better test",
		},
	)
	assert.StringContains(t, "Updated memory", result)
	m, e := s.Store().GetMemory(1)
	assert.FatalOnError(t, e)
	assert.String(t, "test memory", m.Name)
	assert.String(t, "original", m.Content)
	assert.String(t, "a better test", m.Description)
}

func TestUpdateMemoryRenameListsRewrittenCitations(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "pace",
			constant.Content:     "original",
			constant.Description: "a test",
		},
	)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "citing",
			constant.Content:     "See `memory://default/pace`.",
			constant.Description: "a test",
		},
	)
	result := s.MustCallTool(
		constant.UpdateMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.MemoryName:       "tempo",
		},
	)
	assert.StringContains(t, "Rewrote citations in: 2 citing", result)
}

func TestUpdateMemoryRequiresIdentifier(t *testing.T) {
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
