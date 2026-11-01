package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerExec() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ExecInPod,
			mcp.WithDescription(
				"Execute a command in a pod. Command must be an array of strings - no shell interpretation",
			),
			mcp.WithString("name", mcp.Required(), mcp.Description("Pod name")),
			mcp.WithString(
				"namespace",
				mcp.Required(),
				mcp.Description("Namespace"),
			),
			mcp.WithArray(
				"command",
				mcp.Required(),
				mcp.Description(
					"Command as array of strings (e.g. [\"ls\", \"-la\"])",
				),
			),
			mcp.WithString("container", mcp.Description("Container name")),
			mcp.WithNumber(
				"timeout",
				mcp.Description("Timeout in milliseconds (default: 60000)"),
			),
		),
		mcp.NewTypedToolHandler(s.ExecInPod),
	)
}
