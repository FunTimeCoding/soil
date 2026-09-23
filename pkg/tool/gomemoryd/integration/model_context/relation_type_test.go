package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
	"testing"
)

func TestRelateMemoriesWithType(t *testing.T) {
	s := model_context_tester.New(t)
	s.RelatedPair()
	result := s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
			constant.Type:             "affinity",
		},
	)
	assert.StringContains(t, "affinity", result)
	types := s.RelatedTypes()
	assert.Count(t, 1, types)
	assert.String(t, "affinity", types[0])
}

func TestRelateMemoriesRejectsUnknownType(t *testing.T) {
	s := model_context_tester.New(t)
	s.RelatedPair()
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
	s.RelatedPair()
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	types := s.RelatedTypes()
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
	types = s.RelatedTypes()
	assert.Count(t, 1, types)
	assert.String(t, "informs", types[0])
}

func TestUnrelateMemoriesRemovesEdge(t *testing.T) {
	s := model_context_tester.New(t)
	s.RelatedPair()
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
	types := s.RelatedTypes()
	assert.Count(t, 0, types)
}

func TestUnrelateMemoriesIsDirectional(t *testing.T) {
	s := model_context_tester.New(t)
	s.RelatedPair()
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
	types := s.RelatedTypes()
	assert.Count(t, 1, types)
}

func TestRelateMemoriesUntypedKeepsExistingType(t *testing.T) {
	s := model_context_tester.New(t)
	s.RelatedPair()
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
	types := s.RelatedTypes()
	assert.Count(t, 1, types)
	assert.String(t, "grounds", types[0])
}
