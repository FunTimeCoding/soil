package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerAlerts() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAlerts,
			mcp.WithDescription(
				"List alerts from the active Alertmanager instance. By default returns all active, silenced, inhibited, and unprocessed alerts.",
			),
			mcp.WithString(
				"filter",
				mcp.Description(
					"Comma-separated label matchers (e.g. alertname=\"FanFailure\",severity=\"warning\")",
				),
			),
			mcp.WithString(
				"active",
				mcp.Description(
					"Include active alerts (true/false, default true)",
				),
			),
			mcp.WithString(
				"silenced",
				mcp.Description(
					"Include silenced alerts (true/false, default true)",
				),
			),
			mcp.WithString(
				"inhibited",
				mcp.Description(
					"Include inhibited alerts (true/false, default true)",
				),
			),
			mcp.WithString(
				"unprocessed",
				mcp.Description(
					"Include unprocessed alerts (true/false, default true)",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description("Regex to filter alerts by receiver"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of alerts to return"),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of alerts to skip"),
			),
		),
		mcp.NewTypedToolHandler(s.listAlerts),
	)
}
