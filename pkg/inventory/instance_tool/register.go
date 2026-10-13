package instance_tool

import (
	"github.com/funtimecoding/soil/pkg/inventory/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (t *Tool[T, P]) Register(s *server.MCPServer) {
	if len(t.inventory.Instances) < 2 {
		return
	}

	s.AddTool(
		mcp.NewTool(
			constant.ListInstances,
			mcp.WithDescription(
				"List the configured instances. Shows which instance is currently active.",
			),
		),
		mcp.NewTypedToolHandler(t.ListInstances),
	)
	s.AddTool(
		mcp.NewTool(
			constant.UseInstance,
			mcp.WithDescription(
				"Set the active instance for this session. Required before any other tool when more than one is configured.",
			),
			mcp.WithString(
				constant.InstanceParameter,
				mcp.Required(),
				mcp.Description("Instance name from list_instances"),
			),
		),
		mcp.NewTypedToolHandler(t.UseInstance),
	)
}
