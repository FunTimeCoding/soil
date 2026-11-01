package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMonitoring() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RunMonitoring,
			mcp.WithDescription("Run topic monitoring immediately"),
		),
		s.Monitoring,
	)
}
