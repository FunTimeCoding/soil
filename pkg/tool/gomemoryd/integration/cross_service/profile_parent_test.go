//go:build local

package cross_service

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/cross_service_tester"
	"slices"
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
	var profile profileResult
	assert.FatalOnError(t, json.Unmarshal([]byte(raw), &profile))
	indexIDs := make([]int64, len(profile.Index))

	for i, m := range profile.Index {
		indexIDs[i] = m.Identifier
	}

	assert.True(t, slices.Contains(indexIDs, int64(parentIdentifier)))

	for _, m := range profile.Index {
		if m.Identifier == int64(parentIdentifier) {
			assert.Count(t, 3, m.Children)

			for _, name := range childNames {
				assert.True(t, slices.Contains(m.Children, name))
			}

			return
		}
	}

	t.Fatal("parent not found in index")
}
