package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSearch() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchChannels,
			mcp.WithDescription(
				"Search for channels by name. Returns channels you have access to.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Channel name to search for"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchChannels),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchMessages,
			mcp.WithDescription(
				"Search messages across all channels. Supports Mattermost search syntax: from:username, in:channelname, before:2026-04-22, after:2026-04-20. Plain text matches posts containing those words.",
			),
			mcp.WithString(
				"terms",
				mcp.Required(),
				mcp.Description("Search terms"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchMessages),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchUsers,
			mcp.WithDescription(
				"Search users by username, first name, last name, or nickname.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search term"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchUsers),
	)
}
