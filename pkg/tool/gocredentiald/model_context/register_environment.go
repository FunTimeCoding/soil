package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerEnvironment() {
	s.server.AddTool(
		mcp.NewTool(
			constant.LoadGroup,
			mcp.WithDescription(
				"Load an environment variable group by name from the Environment group: custom fields returned as NAME=value pairs, revealed.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Environment group entry title"),
			),
		),
		mcp.NewTypedToolHandler(s.loadGroup),
	)
}
