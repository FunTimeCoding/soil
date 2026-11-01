package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTriage() {
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateIssue,
			mcp.WithDescription(
				"Update the status or assignee of a Sentry issue",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString(
				"status",
				mcp.Description(
					"New status (e.g. resolved, ignored, unresolved)",
				),
			),
			mcp.WithString(
				"assignedTo",
				mcp.Description("Username or team to assign to"),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateIssue),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteIssue,
			mcp.WithDescription("Delete a Sentry issue"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteIssue),
	)
}
