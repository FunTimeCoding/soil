package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTags() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Tag,
			mcp.WithDescription(
				"Set, get, or remove a source type tag for a path prefix. Pass source_type to set, empty string to remove, omit to get.",
			),
			mcp.WithString(
				constant.Path,
				mcp.Required(),
				mcp.Description("Path prefix to tag"),
			),
			mcp.WithString(
				constant.Collection,
				mcp.Description("Scope tag to a collection (omit for global)"),
			),
			mcp.WithString(
				constant.SourceType,
				mcp.Description(
					"Source type to set (omit to get, empty string to remove)",
				),
			),
		),
		s.tag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription("List all configured source type tags."),
		),
		s.listTags,
	)
}
