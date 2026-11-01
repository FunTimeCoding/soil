package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerReading() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Profile,
			mcp.WithDescription(
				"Session start call. Returns three tiers: (1) full content of always-tagged memories, (2) full content of topic-matched memories, (3) brief index of all other active memories (id, description, tags).",
			),
			mcp.WithString(
				constant.Topic,
				mcp.Description(
					"Session topic or context for relevance matching",
				),
			),
			mcp.WithString(
				constant.Scope,
				mcp.Description(
					"Memory scope to profile. Empty for the default scope.",
				),
			),
		),
		s.profile,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMemories,
			mcp.WithDescription(
				"List all active memories. Returns id, description, type, tags, updated_at. Optional type and tag filters.",
			),
			mcp.WithString(
				constant.Type,
				mcp.Description("Filter by memory type"),
			),
			mcp.WithString(constant.Tag, mcp.Description("Filter by tag")),
			mcp.WithString(
				constant.Scope,
				mcp.Description(
					"Memory scope. Empty for the default scope, all crosses scopes.",
				),
			),
		),
		s.list,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchMemories,
			mcp.WithDescription(
				"Search memories by keyword. Returns matching memories ranked by relevance. Optional type and tag filters.",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Required(),
				mcp.Description("Search query"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Description("Maximum results (default 10)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description("Filter by memory type"),
			),
			mcp.WithString(constant.Tag, mcp.Description("Filter by tag")),
			mcp.WithString(
				constant.Scope,
				mcp.Description(
					"Memory scope. Empty for the default scope, all crosses scopes.",
				),
			),
			mcp.WithBoolean(
				constant.Detail,
				mcp.Description(
					"Return full records with timestamps, scope, and type (default false)",
				),
			),
		),
		s.search,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CheckReferences,
			mcp.WithDescription(constant.CheckReferencesSummary),
		),
		s.checkReferences,
	)
}
