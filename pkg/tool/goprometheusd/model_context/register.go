package model_context

import (
	libraryConstant "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	if len(s.service.Instances()) > 1 {
		s.server.AddTool(
			mcp.NewTool(
				constant.ListInstances,
				mcp.WithDescription(
					"List all configured Prometheus instances. Shows which instance is currently active.",
				),
			),
			mcp.NewTypedToolHandler(s.listInstances),
		)
		s.server.AddTool(
			mcp.NewTool(
				constant.UseInstance,
				mcp.WithDescription(
					"Set the active Prometheus instance for this session. Required before using any query tool.",
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
			constant.Query,
			mcp.WithDescription(
				"Execute an instant PromQL query against the active instance.",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("PromQL query expression"),
			),
		),
		mcp.NewTypedToolHandler(s.query),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.QueryRange,
			mcp.WithDescription(
				"Execute a range PromQL query. Returns time series data.",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("PromQL query expression"),
			),
			mcp.WithString(
				"start",
				mcp.Required(),
				mcp.Description("Start time (RFC3339 or duration like 1h)"),
			),
			mcp.WithString(
				"end",
				mcp.Description("End time (RFC3339, default now)"),
			),
			mcp.WithNumber(
				"step",
				mcp.Description("Query step in seconds (default 60)"),
			),
		),
		mcp.NewTypedToolHandler(s.queryRange),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Series,
			mcp.WithDescription(
				"Show TSDB statistics for the active instance.",
			),
		),
		mcp.NewTypedToolHandler(s.series),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.LabelValues,
			mcp.WithDescription(
				"List values for a label name, optionally filtered by series matchers.",
			),
			mcp.WithString(
				libraryConstant.LabelKey,
				mcp.Required(),
				mcp.Description("Label name"),
			),
			mcp.WithString(
				"matches",
				mcp.Description("Comma-separated series matchers"),
			),
			mcp.WithString(
				"since",
				mcp.Description("Lookback duration (e.g. 1h, 24h)"),
			),
		),
		mcp.NewTypedToolHandler(s.labelValues),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.LabelNames,
			mcp.WithDescription(
				"List all label names, optionally filtered by series matchers.",
			),
			mcp.WithString(
				"matches",
				mcp.Description("Comma-separated series matchers"),
			),
			mcp.WithString(
				"since",
				mcp.Description("Lookback duration (e.g. 1h, 24h)"),
			),
		),
		mcp.NewTypedToolHandler(s.labelNames),
	)
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
