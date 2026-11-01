package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerAcrossTypes() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Search,
			mcp.WithDescription(
				"Search links, lists, and tags. Optionally filter by type.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search query string."),
			),
			mcp.WithString(
				"type",
				mcp.Description("Entity type: link, list, tag. Omit for all."),
			),
		),
		mcp.NewTypedToolHandler(s.Search),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Delete,
			mcp.WithDescription(
				"Delete a resource. Type: link, list, tag, note, or branch (deletes list and all its links).",
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Resource type: link, list, tag, note, branch."),
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Resource ID to delete."),
			),
		),
		mcp.NewTypedToolHandler(s.Delete),
	)
}
