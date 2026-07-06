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

func TestProfileHidesNoIndexMemoriesFromIndex(t *testing.T) {
	s := cross_service_tester.New(t)
	result := s.Gomemoryd.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling pattern",
			constant.Content:     "Use captureFail for MCP handlers and captureDetail for per-service wrappers.",
			constant.Description: "Error handling depth leaf",
		},
	)
	var identifier int
	_, e := fmt.Sscanf(result, "Created memory %d", &identifier)
	assert.FatalOnError(t, e)
	s.Gomemoryd.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:              constant.NoIndexTag,
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
	target := int64(identifier)
	indexIDs := make([]int64, len(profile.Index))

	for i, m := range profile.Index {
		indexIDs[i] = m.Identifier
	}

	assert.False(t, slices.Contains(indexIDs, target))
	assert.True(t, slices.Contains(collectIDs(profile.Relevant), target))
}
