package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"slices"
	"testing"
)

func TestDisarmedDaemonRegistersNoWriteTool(t *testing.T) {
	names := base.New(t).Tools().ToolNames()
	assert.Contains(t, []string{"get"}, names)

	for _, tool := range base.WriteTools() {
		assert.False(t, slices.Contains(names, tool))
	}
}

func TestArmedDaemonRegistersWriteTools(t *testing.T) {
	assert.Contains(
		t,
		base.WriteTools(),
		base.New(t, service_tester.WithArmed()).Tools().ToolNames(),
	)
}
