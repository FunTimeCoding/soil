package instance_tool

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/argument"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/instance_response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool[T, P]) ListInstances(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListInstances,
) (*mcp.CallToolResult, error) {
	active, e := t.ActiveName(x)

	if e != nil {
		return response.Answer(t.reporter, e)
	}

	result := make([]*instance_response.Response, 0, len(t.inventory.Instances))

	for k := range t.inventory.Instances {
		i := P(&t.inventory.Instances[k])
		result = append(
			result,
			instance_response.New(
				i.InstanceName(),
				i.InstanceHost(),
				i.InstanceName() == active,
			),
		)
	}

	return response.SuccessAny(result)
}
