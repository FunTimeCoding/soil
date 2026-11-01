package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDirectory() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetUsers,
			mcp.WithDescription("List users in the Mattermost workspace"),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max users to return (default 100)"),
				mcp.DefaultNumber(100),
			),
			mcp.WithNumber(
				"page",
				mcp.Description("Page number starting from 0"),
				mcp.DefaultNumber(0),
			),
		),
		mcp.NewTypedToolHandler(s.GetUsers),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetUserProfile,
			mcp.WithDescription("Get profile information for a user"),
			mcp.WithString(
				"user_id",
				mcp.Required(),
				mcp.Description("User ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetUserProfile),
	)
}
