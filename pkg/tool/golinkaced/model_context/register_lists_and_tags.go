package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerListsAndTags() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLists,
			mcp.WithDescription("List all LinkAce lists."),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListLists),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription("List all LinkAce tags."),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListTags),
	)
}
