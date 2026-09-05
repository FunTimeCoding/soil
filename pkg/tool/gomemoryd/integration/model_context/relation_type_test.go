package model_context

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
	"testing"
)

func relatedPair(
	t *testing.T,
	s *model_context_tester.Tester,
) {
	t.Helper()
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "alpha",
			constant.Content:     "alpha content",
			constant.Description: "alpha description",
		},
	)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "beta",
			constant.Content:     "beta content",
			constant.Description: "beta description",
		},
	)
}

func relatedTypes(
	t *testing.T,
	s *model_context_tester.Tester,
) []string {
	t.Helper()
	raw := s.MustCallTool(
		constant.GetMemory,
		map[string]any{constant.MemoryIdentifier: 1},
	)
	var result struct {
		Related []struct {
			Type string `json:"type"`
		} `json:"related"`
	}
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &result))
	types := make([]string, 0, len(result.Related))

	for _, r := range result.Related {
		types = append(types, r.Type)
	}

	return types
}

func TestRelateMemoriesWithType(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	result := s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
			constant.Type:             "affinity",
		},
	)
	assert.StringContains(t, "affinity", result)
	types := relatedTypes(t, s)
	assert.Count(t, 1, types)
	assert.String(t, "affinity", types[0])
}

func TestRelateMemoriesRejectsUnknownType(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	result := s.MustCallToolError(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
			constant.Type:             "vehicle",
		},
	)
	assert.StringContains(t, "unknown relation type", result)
}

func TestRelateMemoriesRetypesExistingEdge(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	types := relatedTypes(t, s)
	assert.Count(t, 1, types)
	assert.String(t, "", types[0])
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
			constant.Type:             "informs",
		},
	)
	types = relatedTypes(t, s)
	assert.Count(t, 1, types)
	assert.String(t, "informs", types[0])
}

func TestUnrelateMemoriesRemovesEdge(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	result := s.MustCallTool(
		constant.UnrelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	assert.StringContains(t, "Unrelated", result)
	types := relatedTypes(t, s)
	assert.Count(t, 0, types)
}

func TestUnrelateMemoriesIsDirectional(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	result := s.MustCallToolError(
		constant.UnrelateMemories,
		map[string]any{
			constant.SourceIdentifier: 2,
			constant.TargetIdentifier: 1,
		},
	)
	assert.StringContains(t, "reverse", result)
	types := relatedTypes(t, s)
	assert.Count(t, 1, types)
}

func TestRelateMemoriesUntypedKeepsExistingType(t *testing.T) {
	s := model_context_tester.New(t)
	relatedPair(t, s)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
			constant.Type:             "grounds",
		},
	)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	types := relatedTypes(t, s)
	assert.Count(t, 1, types)
	assert.String(t, "grounds", types[0])
}
