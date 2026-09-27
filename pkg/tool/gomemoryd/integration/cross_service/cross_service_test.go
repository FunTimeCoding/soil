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

func TestProfileShowsCompletions(t *testing.T) {
	s := cross_service_tester.New(t)
	s.QueryClient.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/1",
			goquerydConstant.Body:       "built the search pipeline",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.MemoryClient.MustCallTool(constant.Profile, map[string]any{})
	assert.StringContains(t, "test-session", result)
	assert.StringContains(t, "built the search pipeline", result)
}

func TestProfileShowsMultipleCompletionsPerSession(t *testing.T) {
	s := cross_service_tester.New(t)
	s.QueryClient.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/1",
			goquerydConstant.Body:       "built the API",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	s.QueryClient.MustCallTool(
		goquerydConstant.Push,
		map[string]any{
			goquerydConstant.Collection: "completions",
			goquerydConstant.Path:       "test-session/2",
			goquerydConstant.Body:       "wrote the tests",
			goquerydConstant.Metadata: map[string]string{
				"source_type":  "session-completion",
				"session_name": "test-session",
			},
		},
	)
	result := s.MemoryClient.MustCallTool(constant.Profile, map[string]any{})
	assert.StringContains(t, "built the API", result)
	assert.StringContains(t, "wrote the tests", result)
}

func TestProfileCompletionsEmptyWhenNone(t *testing.T) {
	assert.StringNotContains(
		t,
		"completions",
		cross_service_tester.New(t).MemoryClient.MustCallTool(
			constant.Profile,
			map[string]any{},
		),
	)
}

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

func TestProfileHidesNoIndexMemoriesFromIndex(t *testing.T) {
	s := cross_service_tester.New(t)
	result := s.MemoryClient.MustCallTool(
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
	s.MemoryClient.MustCallTool(
		constant.TagMemory,
		map[string]any{
			constant.MemoryIdentifier: identifier,
			constant.Add:              constant.NoIndexTag,
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
	assert.StringNotContains(
		t,
		fmt.Sprintf("%d error handling pattern", identifier),
		fixture.Section(raw, constant.IndexSectionHeading),
	)
	assert.StringContains(
		t,
		mark,
		fixture.Section(raw, constant.RelevantSectionHeading),
	)
}

func TestProfileCollapsesChildrenUnderParent(t *testing.T) {
	s := cross_service_tester.New(t)
	parentResult := s.MemoryClient.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling",
			constant.Content:     "Error handling patterns and conventions.",
			constant.Description: "Error handling parent",
		},
	)
	var parentIdentifier int
	_, e := fmt.Sscanf(parentResult, "Created memory %d", &parentIdentifier)
	assert.FatalOnError(t, e)
	childNames := []string{"captureFail", "captureDetail", "clientError"}

	for _, name := range childNames {
		s.MemoryClient.MustCallTool(
			constant.SaveMemory,
			map[string]any{
				constant.MemoryName:       name,
				constant.Content:          fmt.Sprintf("%s content", name),
				constant.Description:      fmt.Sprintf("%s description", name),
				constant.ParentIdentifier: parentIdentifier,
			},
		)
	}

	raw := s.MemoryClient.MustCallTool(constant.Profile, map[string]any{})
	index := fixture.Section(raw, constant.IndexSectionHeading)
	assert.StringContains(t, fmt.Sprintf("%d ", parentIdentifier), index)

	for _, name := range childNames {
		assert.StringContains(t, name, index)
	}
}

func TestProfileRelevantTierUsesHybridSearch(t *testing.T) {
	s := cross_service_tester.New(t)
	s.MemoryClient.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling pattern",
			constant.Content:     "Use captureFail for MCP handlers and captureDetail for per-service wrappers.",
			constant.Description: "MCP error handling conventions",
		},
	)
	s.QueryClient.MustCallTool(goquerydConstant.Embed, map[string]any{})
	result := s.MemoryClient.MustCallTool(
		constant.Profile,
		map[string]any{
			constant.Topic: "how to handle errors in model context tools",
		},
	)
	assert.StringContains(t, "captureFail", result)
}
