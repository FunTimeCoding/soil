package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Query,
			mcp.WithDescription(
				"Query recent usage events. Filter by tool name, surface (mcp, cli, rest, web), actor, and time range. Returns individual events, newest first.",
			),
			mcp.WithString(
				constant.Tool,
				mcp.Description("Filter by tool name."),
			),
			mcp.WithString(
				constant.Surface,
				mcp.Description(
					"Filter by surface: mcp, cli, rest, web.",
				),
			),
			mcp.WithString(
				constant.Actor,
				mcp.Description("Filter by actor name."),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"Filter by kind: baseline or domain.",
				),
			),
			mcp.WithString(
				constant.Since,
				mcp.Description("RFC3339 timestamp lower bound."),
			),
			mcp.WithString(
				constant.Until,
				mcp.Description("RFC3339 timestamp upper bound."),
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description(
					"Maximum events to return. Default: 50.",
				),
			),
		),
		s.query,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Summary,
			mcp.WithDescription(
				"Aggregated usage summary - counts per tool, optionally broken down by surface. The heatmap.",
			),
			mcp.WithString(
				constant.Since,
				mcp.Description("RFC3339 timestamp lower bound."),
			),
			mcp.WithString(
				constant.Until,
				mcp.Description("RFC3339 timestamp upper bound."),
			),
			mcp.WithString(
				constant.GroupBy,
				mcp.Description(
					"Group by: tool (default), surface, or kind.",
				),
			),
		),
		s.summary,
	)
}
