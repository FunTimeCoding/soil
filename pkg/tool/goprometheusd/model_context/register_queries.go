package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerQueries() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Query,
			mcp.WithDescription(
				"Execute an instant PromQL query against the active instance.",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
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
				generativeConstant.ParameterQuery,
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
}
