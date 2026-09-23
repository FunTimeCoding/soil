//go:build local

package cross_service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/cross_service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/fixture"
	"testing"
)

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
