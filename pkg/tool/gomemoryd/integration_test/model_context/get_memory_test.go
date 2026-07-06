//go:build local

package model_context

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/model_context_tester"
	"testing"
)

type getMemoryResult struct {
	Identifier int64          `json:"identifier"`
	Name       string         `json:"name"`
	Related    []relatedEntry `json:"related"`
}

type relatedEntry struct {
	Identifier  int64    `json:"identifier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func TestGetMemoryIncludesRelated(t *testing.T) {
	s := model_context_tester.New(t)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling",
			constant.Content:     "error handling content",
			constant.Description: "error handling description",
		},
	)
	s.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "deployment",
			constant.Content:     "deployment content",
			constant.Description: "deployment description",
		},
	)
	s.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: 2,
			constant.Add:              "deploy,no-index",
		},
	)
	s.MustCallTool(
		constant.RelateMemories,
		map[string]any{
			constant.SourceIdentifier: 1,
			constant.TargetIdentifier: 2,
		},
	)
	raw := s.MustCallTool(
		constant.GetMemory,
		map[string]any{constant.MemoryIdentifier: 1},
	)
	var result getMemoryResult
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &result))
	assert.String(t, "error handling", result.Name)
	assert.Count(t, 1, result.Related)
	assert.Integer64(t, 2, result.Related[0].Identifier)
	assert.String(t, "deployment", result.Related[0].Name)
	assert.String(
		t,
		"deployment description",
		result.Related[0].Description,
	)
	assert.Count(t, 2, result.Related[0].Tags)
}
