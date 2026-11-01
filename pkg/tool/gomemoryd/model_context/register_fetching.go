package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerFetching() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMemory,
			mcp.WithDescription(
				"Get one or several memories. Returns content, description, tags, metadata, and related memories (identifier, name, description, type, tags per neighbor). Optionally includes version history (single fetch only).",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Description("Memory ID (single fetch)"),
			),
			mcp.WithString(
				constant.MemoryIdentifiers,
				mcp.Description(
					"Comma-separated memory IDs - returns an array, one entry per memory with its relations",
				),
			),
			mcp.WithBoolean(
				constant.IncludeHistory,
				mcp.Description("Include version history (default false)"),
			),
			mcp.WithBoolean(
				constant.Detail,
				mcp.Description(
					"Return the full record with timestamps, scope, and provenance (default false)",
				),
			),
		),
		s.get,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMemoryGroup,
			mcp.WithDescription(
				"Get a parent memory with the full content of all its children in one call, ordered by ordinal then recency. The group fetch for document-sourced memory clusters.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Parent memory ID"),
			),
			mcp.WithBoolean(
				constant.Detail,
				mcp.Description(
					"Return full records with timestamps, scope, and provenance (default false)",
				),
			),
		),
		s.getGroup,
	)
}
