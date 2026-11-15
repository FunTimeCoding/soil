package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/model_context_tester"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"slices"
	"testing"
)

func TestDisarmedDaemonRegistersNoWriteTool(t *testing.T) {
	c, _ := service_tester.New(t).Mounted(false)
	names := model_context_tester.ToolNames(c)
	assert.Contains(t, []string{"get"}, names)

	for _, tool := range model_context_tester.WriteTools() {
		assert.False(t, slices.Contains(names, tool))
	}
}

func TestArmedDaemonRegistersWriteTools(t *testing.T) {
	c, _ := service_tester.New(t).Mounted(true)
	assert.Contains(
		t,
		model_context_tester.WriteTools(),
		model_context_tester.ToolNames(c),
	)
}
