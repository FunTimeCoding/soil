package model_context

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSeries() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Series,
			mcp.WithDescription("Show TSDB statistics for the active instance."),
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
				soil.LabelKey,
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
}
