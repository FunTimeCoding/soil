package instance_tool

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/inventory/instance_tool/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (t *Tool[T, P]) UseInstance(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UseInstance,
) (*mcp.CallToolResult, error) {
	if a.Instance == "" {
		return response.Fail("instance is required")
	}

	if _, okay := t.inventory.Find(a.Instance); !okay {
		return response.Fail("unknown instance: %s", a.Instance)
	}

	if e := t.selection.Select(x, a.Instance); e != nil {
		return response.Answer(t.reporter, e)
	}

	return response.Success("active instance set to %s", a.Instance)
}
