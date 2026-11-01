package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRules() {
	s.server.AddTool(
		mcp.NewTool(
			constant.QueryRules,
			mcp.WithDescription(
				"List rules with evaluation health, per-alert state, and activeAt timestamps. Defaults to alerting rules only. Useful for debugging why alerts are or aren't firing.",
			),
			mcp.WithString(
				"name",
				mcp.Description("Filter by rule name (substring match)"),
			),
			mcp.WithString(
				"group",
				mcp.Description("Filter by rule group name (substring match)"),
			),
			mcp.WithString(
				"type",
				mcp.Description("Rule type: alert (default), record, or all"),
			),
			mcp.WithString(
				"state",
				mcp.Description("Filter by state: pending, firing, inactive"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of rules to return"),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of rules to skip"),
			),
		),
		mcp.NewTypedToolHandler(s.queryRules),
	)
}
