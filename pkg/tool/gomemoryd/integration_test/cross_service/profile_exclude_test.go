//go:build local

package cross_service

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/cross_service_tester"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"slices"
	"testing"
)

type profileResult struct {
	Always   []profileMemory  `json:"always"`
	Relevant []profileMemory  `json:"relevant"`
	Index    []profileSummary `json:"index"`
}

type profileMemory struct {
	Identifier int64  `json:"identifier"`
	Name       string `json:"name"`
}

type profileSummary struct {
	Identifier int64    `json:"identifier"`
	Name       string   `json:"name"`
	Children   []string `json:"children,omitempty"`
}

func collectIDs(memories []profileMemory) []int64 {
	result := make([]int64, len(memories))

	for i, m := range memories {
		result[i] = m.Identifier
	}

	return result
}

func TestProfileExcludesAlwaysMemoriesFromRelevant(t *testing.T) {
	s := cross_service_tester.New(t)
	result := s.Gomemoryd.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling pattern",
			constant.Content:     "Use captureFail for MCP handlers and captureDetail for per-service wrappers.",
			constant.Description: "MCP error handling conventions",
		},
	)
	var identifier int
	_, e := fmt.Sscanf(result, "Created memory %d", &identifier)
	assert.FatalOnError(t, e)
	s.Gomemoryd.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:             "always",
		},
	)
	s.Gomemoryd.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "deployment pipeline",
			constant.Content:     "Deployment orchestration uses ArgoCD with kustomize overlays.",
			constant.Description: "Deployment pipeline conventions",
		},
	)
	s.Goqueryd.MustCallTool(
		goquerydConstant.Embed,
		map[string]any{},
	)
	raw := s.Gomemoryd.MustCallTool(
		constant.Profile,
		map[string]any{
			constant.Topic: "error handling patterns in MCP services",
		},
	)
	var profile profileResult
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &profile))
	alwaysIDs := collectIDs(profile.Always)
	relevantIDs := collectIDs(profile.Relevant)
	target := int64(identifier)
	assert.True(t, slices.Contains(alwaysIDs, target))
	assert.False(t, slices.Contains(relevantIDs, target))
}
