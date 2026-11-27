package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerUsers() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchUsers,
			mcp.WithDescription(
				"Search people in the organization whose name or email starts with the query",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Start of a display name or email address"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of people (default 20)"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchUsers),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetUserProfile,
			mcp.WithDescription(
				"Get a person's profile by display name, email address or identifier",
			),
			mcp.WithString(
				"user",
				mcp.Required(),
				mcp.Description("Display name, email address or identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.GetUserProfile),
	)
}
