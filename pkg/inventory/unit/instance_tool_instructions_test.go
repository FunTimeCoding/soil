package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	"github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool"
	"testing"
)

func TestInstanceToolInstructionsFollowTheCount(t *testing.T) {
	i := identity.New("probe", "probe tool", "probe").WithInstructions(
		"Probe bridge.",
		paragraph.New(
			constant.MultiInstance,
			constant.MultiInstanceInstruction,
		),
	)
	assert.String(t, "Probe bridge.", instance_tool.Instructions(i, 1))
	assert.String(
		t,
		"Probe bridge. Multi-instance - call list_instances and use_instance before querying.",
		instance_tool.Instructions(i, 2),
	)
}
