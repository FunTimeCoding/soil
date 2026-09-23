//go:build local

package cross_service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/cross_service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/fixture"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestProfileExcludesAlwaysMemoriesFromRelevant(t *testing.T) {
	s := cross_service_tester.New(t)
	result := s.MemoryClient.MustCallTool(
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
	s.MemoryClient.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:              "always",
		},
	)
	s.MemoryClient.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "deployment pipeline",
			constant.Content:     "Deployment orchestration uses ArgoCD with kustomize overlays.",
			constant.Description: "Deployment pipeline conventions",
		},
	)
	s.QueryClient.MustCallTool(goquerydConstant.Embed, map[string]any{})
	raw := s.MemoryClient.MustCallTool(
		constant.Profile,
		map[string]any{
			constant.Topic: "error handling patterns in MCP services",
		},
	)
	mark := fmt.Sprintf("(%d)", identifier)
	assert.StringContains(
		t,
		mark,
		fixture.Section(raw, constant.AlwaysSectionHeading),
	)
	assert.StringNotContains(
		t,
		mark,
		fixture.Section(raw, constant.RelevantSectionHeading),
	)
}
