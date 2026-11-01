package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerStatus() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetStatus,
			mcp.WithDescription("Get Alertmanager cluster status and version."),
			mcp.WithBoolean(
				"include_configuration",
				mcp.Description(
					"Include the full alertmanager routing configuration",
				),
			),
		),
		mcp.NewTypedToolHandler(s.getStatus),
	)
}
