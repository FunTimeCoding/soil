package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPortForwards() {
	s.server.AddTool(
		mcp.NewTool(
			constant.PortForward,
			mcp.WithDescription(
				"Forward a local port to a port on a Kubernetes resource",
			),
			mcp.WithString(
				"resourceType",
				mcp.Required(),
				mcp.Description("Resource type (e.g. pod, service)"),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Resource name"),
			),
			mcp.WithString(
				"namespace",
				mcp.Description("Namespace (default: default)"),
			),
			mcp.WithNumber(
				"localPort",
				mcp.Required(),
				mcp.Description("Local port"),
			),
			mcp.WithNumber(
				"targetPort",
				mcp.Required(),
				mcp.Description("Target port on the resource"),
			),
		),
		mcp.NewTypedToolHandler(s.PortForward),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.StopPortForward,
			mcp.WithDescription("Stop an active port forward"),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("Port forward ID"),
			),
		),
		mcp.NewTypedToolHandler(s.StopPortForward),
	)
}
