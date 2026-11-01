package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSilences() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSilences,
			mcp.WithDescription(
				"List silences. By default only active silences are shown.",
			),
			mcp.WithBoolean(
				"expired",
				mcp.Description("Include expired silences"),
			),
			mcp.WithString(
				"filter",
				mcp.Description("Filter by rule name (substring match)"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of silences to return"),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of silences to skip"),
			),
		),
		mcp.NewTypedToolHandler(s.listSilences),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateSilence,
			mcp.WithDescription(
				"Create or extend a silence for an alert rule. Validates that the rule exists. If a silence already exists for the rule, it extends it.",
			),
			mcp.WithString(
				"alert",
				mcp.Required(),
				mcp.Description("Alert rule name to silence"),
			),
			mcp.WithString("comment", mcp.Description("Silence comment")),
			mcp.WithString(
				"duration",
				mcp.Description(
					"Silence duration (e.g. 10m, 1h, 24h). Default 10m.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.createSilence),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteSilence,
			mcp.WithDescription("Expire a silence by its identifier."),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("Silence identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteSilence),
	)
}
