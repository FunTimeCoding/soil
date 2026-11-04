package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"slices"
	"testing"
)

func TestADisarmedDaemonRegistersNoWriteTool(t *testing.T) {
	c, _ := service_tester.New(t).Mounted(false)
	names := toolNames(c)
	assert.Contains(t, []string{"get"}, names)

	for _, tool := range writeTools() {
		assert.False(t, slices.Contains(names, tool))
	}
}

func TestAnArmedDaemonRegistersTheWriteTools(t *testing.T) {
	c, _ := service_tester.New(t).Mounted(true)
	assert.Contains(t, writeTools(), toolNames(c))
}
