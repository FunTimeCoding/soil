package model_context

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/fixture"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/model_context_tester"
	"strings"
	"testing"
)

func TestGetMemoryCompactByDefault(t *testing.T) {
	s := model_context_tester.New(t)
	identifier := s.DocumentSourcedMemory()
	result := fixture.Decode(
		t,
		s.MustCallTool(
			constant.GetMemory,
			map[string]any{constant.MemoryIdentifier: identifier},
		),
	)
	assert.String(t, "Retry", result[constant.MemoryName].(string))
	assert.String(
		t,
		"Document-sourced content.",
		result[constant.Content].(string),
	)

	for _, key := range []string{
		"provenance_file",
		"provenance_hash",
		"created_at",
		"updated_at",
		"is_active",
		constant.Type,
		constant.Scope,
	} {
		if _, found := result[key]; found {
			t.Errorf("compact response carries %s", key)
		}
	}
}

func TestGetMemoryDetailCarriesProvenance(t *testing.T) {
	s := model_context_tester.New(t)
	identifier := s.DocumentSourcedMemory()
	result := fixture.Decode(
		t,
		s.MustCallTool(
			constant.GetMemory,
			map[string]any{
				constant.MemoryIdentifier: identifier,
				constant.Detail:           true,
			},
		),
	)
	assert.String(t, "canon/Example.yaml", result["provenance_file"].(string))
	assert.String(t, "alfa", result[constant.Scope].(string))
	assert.True(t, result["is_active"].(bool))
}

func TestGetMemoryGroupCompactByDefault(t *testing.T) {
	s := model_context_tester.New(t)
	parent := s.EmptyContentParent()
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:       "shard",
			constant.Content:          "shard content",
			constant.Description:      "shard description",
			constant.ParentIdentifier: parent,
		},
	)
	result := fixture.Decode(
		t,
		s.MustCallTool(
			constant.GetMemoryGroup,
			map[string]any{constant.MemoryIdentifier: parent},
		),
	)
	parentResult := result["parent"].(map[string]any)

	if _, found := parentResult["provenance_hash"]; found {
		t.Error("compact group parent carries provenance_hash")
	}

	if _, found := parentResult[constant.Content]; found {
		t.Error("compact group parent carries empty content")
	}

	children := result["children"].([]any)
	assert.Count(t, 1, children)
	child := children[0].(map[string]any)
	assert.String(t, "shard", child[constant.MemoryName].(string))

	if _, found := child["created_at"]; found {
		t.Error("compact group child carries created_at")
	}
}

func TestGetMemoryGroupCarriesRelations(t *testing.T) {
	s := model_context_tester.New(t)
	parent := s.EmptyContentParent()
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:       "shard",
			constant.Content:          "shard content",
			constant.Description:      "shard description",
			constant.ParentIdentifier: parent,
		},
	)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "outside",
			constant.Content:     "outside content",
			constant.Description: "outside description",
		},
	)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: parent,
			constant.TargetIdentifier: 3,
			constant.Type:             "deep-dive",
		},
	)
	result := fixture.Decode(
		t,
		s.MustCallTool(
			constant.GetMemoryGroup,
			map[string]any{constant.MemoryIdentifier: parent},
		),
	)
	relations := result["relations"].([]any)
	assert.Count(t, 1, relations)
	first := relations[0].(map[string]any)
	assert.String(t, "outside", first["target"].(string))
	assert.String(t, "deep-dive", first[constant.Type].(string))
}

func TestGetMemoryGroupDetailCarriesProvenance(t *testing.T) {
	s := model_context_tester.New(t)
	parent := s.DocumentSourcedMemory()
	result := fixture.Decode(
		t,
		s.MustCallTool(
			constant.GetMemoryGroup,
			map[string]any{
				constant.MemoryIdentifier: parent,
				constant.Detail:           true,
			},
		),
	)
	parentResult := result["parent"].(map[string]any)
	assert.String(
		t,
		"canon/Example.yaml",
		parentResult["provenance_file"].(string),
	)
}

func TestSearchMemoriesCompactByDefault(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "retry policy",
			constant.Content:     "retry with backoff",
			constant.Description: "retry description",
		},
	)
	raw := s.MustCallTool(
		constant.SearchMemories,
		map[string]any{constant.Query: "retry"},
	)
	var results []map[string]any
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &results))
	assert.Count(t, 1, results)
	assert.String(t, "retry policy", results[0][constant.MemoryName].(string))

	for _, key := range []string{constant.Type, "updated_at"} {
		if _, found := results[0][key]; found {
			t.Errorf("compact search result carries %s", key)
		}
	}
}

func TestSearchMemoriesDetailCarriesType(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "retry policy",
			constant.Content:     "retry with backoff",
			constant.Description: "retry description",
		},
	)
	raw := s.MustCallTool(
		constant.SearchMemories,
		map[string]any{constant.Query: "retry", constant.Detail: true},
	)
	var results []map[string]any
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &results))
	assert.Count(t, 1, results)
	assert.String(t, "feedback", results[0][constant.Type].(string))
}

func TestGetMemoryManyReturnsArray(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "first",
			constant.Content:     "first content",
			constant.Description: "first description",
		},
	)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "second",
			constant.Content:     "second content",
			constant.Description: "second description",
		},
	)
	raw := s.MustCallTool(
		constant.GetMemory,
		map[string]any{constant.MemoryIdentifiers: "1, 2"},
	)
	var results []map[string]any
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &results))
	assert.Count(t, 2, results)
	assert.String(t, "first", results[0][constant.MemoryName].(string))
	assert.String(t, "second", results[1][constant.MemoryName].(string))

	if _, found := results[0]["created_at"]; found {
		t.Error("compact many-fetch carries created_at")
	}
}

func TestProfileAlwaysTierCompact(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "charlie",
			constant.Content:     "charlie content",
			constant.Description: "charlie description",
		},
	)
	s.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: 1,
			constant.Add:              constant.AlwaysTag,
		},
	)
	always := fixture.Section(
		s.MustCallTool(constant.Profile, map[string]any{}),
		constant.AlwaysSectionHeading,
	)
	assert.Integer(t, 1, strings.Count(always, constant.MemoryHeadingPrefix))
	assert.StringContains(t, "## charlie (1)", always)
	assert.StringContains(t, "charlie content", always)

	for _, key := range []string{"created_at", "is_active", constant.Type} {
		assert.StringNotContains(t, key, always)
	}
}
