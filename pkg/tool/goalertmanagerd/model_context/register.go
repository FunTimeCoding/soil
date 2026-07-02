package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	if len(s.service.Instances()) > 1 {
		s.server.AddTool(
			mcp.NewTool(
				constant.ListInstances,
				mcp.WithDescription(
					"List all configured Alertmanager instances. Shows which instance is currently active.",
				),
			),
			mcp.NewTypedToolHandler(s.listInstances),
		)
		s.server.AddTool(
			mcp.NewTool(
				constant.UseInstance,
				mcp.WithDescription(
					"Set the active Alertmanager instance for this session. Required before using any other tool.",
				),
				mcp.WithString(
					"instance",
					mcp.Required(),
					mcp.Description("Instance name from list_instances"),
				),
			),
			mcp.NewTypedToolHandler(s.useInstance),
		)
	}

	s.server.AddTool(
		mcp.NewTool(
			constant.ListAlerts,
			mcp.WithDescription(
				"List alerts from the active Alertmanager instance. By default returns all active, silenced, inhibited, and unprocessed alerts.",
			),
			mcp.WithString(
				"filter",
				mcp.Description("Comma-separated label matchers (e.g. alertname=\"FanFailure\",severity=\"warning\")"),
			),
			mcp.WithString(
				"active",
				mcp.Description("Include active alerts (true/false, default true)"),
			),
			mcp.WithString(
				"silenced",
				mcp.Description("Include silenced alerts (true/false, default true)"),
			),
			mcp.WithString(
				"inhibited",
				mcp.Description("Include inhibited alerts (true/false, default true)"),
			),
			mcp.WithString(
				"unprocessed",
				mcp.Description("Include unprocessed alerts (true/false, default true)"),
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
			mcp.WithString(
				"comment",
				mcp.Description("Silence comment"),
			),
			mcp.WithString(
				"duration",
				mcp.Description("Silence duration (e.g. 10m, 1h, 24h). Default 10m."),
			),
		),
		mcp.NewTypedToolHandler(s.createSilence),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteSilence,
			mcp.WithDescription(
				"Expire a silence by its identifier.",
			),
			mcp.WithString(
				"id",
				mcp.Required(),
				mcp.Description("Silence identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteSilence),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetStatus,
			mcp.WithDescription(
				"Get Alertmanager cluster status and version.",
			),
			mcp.WithBoolean(
				"include_configuration",
				mcp.Description("Include the full alertmanager routing configuration"),
			),
		),
		mcp.NewTypedToolHandler(s.getStatus),
	)
}
