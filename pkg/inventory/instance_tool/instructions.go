package instance_tool

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/inventory/constant"
)

func Instructions(
	i *identity.Tool,
	instances int,
) string {
	return i.RenderInstructions(
		map[string]bool{constant.MultiInstance: instances > 1},
	)
}
