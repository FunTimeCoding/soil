package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetAlerts,
			mcp.WithDescription("Get alert log entries matching the given name"),
			mcp.WithString(
				parameter.Name,
				mcp.Required(),
				mcp.Description("Alert name to filter by"),
			),
		),
		s.getAlerts,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetRecentAlerts,
			mcp.WithDescription(
				"Get alert log entries within a time range (defaults to last 1 hour)",
			),
			mcp.WithString(
				constant.Start,
				mcp.Description(
					"Start of time range (RFC3339 format) - defaults to 1 hour ago",
				),
			),
			mcp.WithString(
				constant.End,
				mcp.Description(
					"End of time range (RFC3339 format) - defaults to now",
				),
			),
		),
		s.getRecentAlerts,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetTopAlerts,
			mcp.WithDescription(
				"Get top alerts by occurrence count (defaults to top 25 over last 7 days)",
			),
			mcp.WithNumber(
				constant.N,
				mcp.Description("Number of top alerts to return (default 25)"),
			),
			mcp.WithString(
				constant.Start,
				mcp.Description(
					"Start of time range (RFC3339 format) - defaults to 7 days ago",
				),
			),
			mcp.WithString(
				constant.End,
				mcp.Description(
					"End of time range (RFC3339 format) - defaults to now",
				),
			),
		),
		s.getTopAlerts,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetStatus,
			mcp.WithDescription(
				"Get alert log service status including total records and last poll time",
			),
		),
		s.getStatus,
	)
}
