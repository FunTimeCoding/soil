//go:build local

package cross_service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration_test/cross_service_tester"
	goquerydConstant "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestProfileRelevantTierUsesHybridSearch(t *testing.T) {
	s := cross_service_tester.New(t)
	s.Gomemoryd.MustCallTool(
		constant.SaveMemory,
		map[string]any{
			constant.MemoryName:  "error handling pattern",
			constant.Content:     "Use captureFail for MCP handlers and captureDetail for per-service wrappers.",
			constant.Description: "MCP error handling conventions",
		},
	)
	s.Goqueryd.MustCallTool(
		goquerydConstant.Embed,
		map[string]any{},
	)
	result := s.Gomemoryd.MustCallTool(
		constant.Profile,
		map[string]any{
			constant.Topic: "how to handle errors in model context tools",
		},
	)
	assert.StringContains(t, "captureFail", result)
}
