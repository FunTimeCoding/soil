package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTags() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription(
				"List all tags with memory counts. Returns each tag and how many active memories carry it, ordered by count.",
			),
		),
		s.listTags,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TagMemory,
			mcp.WithDescription(
				"Add, remove, or replace tags on a memory. At least one of add, remove, or replace_all is required. Tags are comma-separated. Special tags: 'always' loads the memory in full on every profile call, 'no-index' hides it from the profile index (still reachable via topic matching, search, and relations).",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Memory ID"),
			),
			mcp.WithString(
				constant.Add,
				mcp.Description("Comma-separated tags to add"),
			),
			mcp.WithString(
				constant.Remove,
				mcp.Description("Comma-separated tags to remove"),
			),
			mcp.WithString(
				constant.ReplaceAll,
				mcp.Description(
					"Comma-separated tags to replace all existing tags with",
				),
			),
		),
		s.tag,
	)
}
