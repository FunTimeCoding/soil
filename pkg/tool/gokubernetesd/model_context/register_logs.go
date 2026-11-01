package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerLogs() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Logs,
			mcp.WithDescription(
				"Get logs from a pod or resource that owns pods (e.g. deployment/sentry, job/migrate)",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description(
					"Pod name or resource/name (e.g. sentry-abc123, deployment/sentry, job/migrate)",
				),
			),
			mcp.WithString(
				"namespace",
				mcp.Required(),
				mcp.Description("Namespace"),
			),
			mcp.WithString(
				"container",
				mcp.Description(
					"Container name (required for multi-container pods)",
				),
			),
			mcp.WithNumber(
				"tail",
				mcp.Description("Number of lines from the end"),
			),
			mcp.WithString(
				"since",
				mcp.Description("Relative time (e.g. 5m, 1h)"),
			),
			mcp.WithBoolean(
				"previous",
				mcp.Description("Logs from previously terminated container"),
			),
			mcp.WithBoolean(
				"timestamps",
				mcp.Description("Include timestamps"),
			),
			mcp.WithBoolean(
				"all",
				mcp.Description(
					"Get logs from all matching pods (default: error with list if multiple)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Logs),
	)
}
