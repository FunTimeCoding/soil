package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SaveMemory,
			mcp.WithDescription(
				"Create a new memory. Returns the created memory ID.",
			),
			mcp.WithString(
				constant.MemoryName,
				mcp.Required(),
				mcp.Description("Short name for identification (e.g. 'retry policy', 'careful work')"),
			),
			mcp.WithString(
				constant.Content,
				mcp.Required(),
				mcp.Description("The memory content"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Required(),
				mcp.Description("One-line description for index and search"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description("Memory type: user, feedback, project, reference (default: feedback)"),
			),
			mcp.WithString(
				constant.Source,
				mcp.Description("Session name of the caller (optional, for version attribution)"),
			),
			mcp.WithNumber(
				constant.ParentIdentifier,
				mcp.Description("Parent memory ID (optional, for hierarchy)"),
			),
		),
		s.create,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateMemory,
			mcp.WithDescription(
				"Update an existing memory. Records the previous version in history. Tags are preserved.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("ID of the memory to update"),
			),
			mcp.WithString(
				constant.MemoryName,
				mcp.Required(),
				mcp.Description("Short name for identification"),
			),
			mcp.WithString(
				constant.Content,
				mcp.Required(),
				mcp.Description("The updated memory content"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Required(),
				mcp.Description("One-line description for index and search"),
			),
			mcp.WithString(
				constant.Source,
				mcp.Description("Session name of the caller (optional, for version attribution)"),
			),
		),
		s.update,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Profile,
			mcp.WithDescription(
				"Session start call. Returns three tiers: (1) full content of always-tagged memories, (2) full content of topic-matched memories, (3) brief index of all other active memories (id, description, tags).",
			),
			mcp.WithString(
				constant.Topic,
				mcp.Description("Session topic or context for relevance matching"),
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
			mcp.WithString(
				constant.Tag,
				mcp.Description("Filter by tag"),
			),
		),
		s.list,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMemory,
			mcp.WithDescription(
				"Get a memory by ID. Returns full content, description, type, tags, and related memories (identifier, name, description, tags per neighbor). Optionally includes version history.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Memory ID"),
			),
			mcp.WithBoolean(
				constant.IncludeHistory,
				mcp.Description("Include version history (default false)"),
			),
		),
		s.get,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ForgetMemory,
			mcp.WithDescription(
				"Soft-delete a memory. The memory is marked inactive and a version record is created. Content is preserved in history.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Memory ID to forget"),
			),
		),
		s.forget,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchMemories,
			mcp.WithDescription(
				"Search memories by keyword. Returns matching memories ranked by relevance. Optional type and tag filters.",
			),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("Search query"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Description("Maximum results (default 10)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description("Filter by memory type"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Description("Filter by tag"),
			),
		),
		s.search,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RelateMemories,
			mcp.WithDescription(
				"Create a relation between two memories. Relations are bidirectional - get_memory on either shows the other under related.",
			),
			mcp.WithNumber(
				constant.SourceIdentifier,
				mcp.Required(),
				mcp.Description("First memory ID"),
			),
			mcp.WithNumber(
				constant.TargetIdentifier,
				mcp.Required(),
				mcp.Description("Second memory ID"),
			),
		),
		s.relate,
	)
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
				mcp.Description("Comma-separated tags to replace all existing tags with"),
			),
		),
		s.tag,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetParent,
			mcp.WithDescription(
				"Set or unset the parent of a memory. Pass parent_id to set, omit to unset. Children are collapsed under their parent in the profile index.",
			),
			mcp.WithNumber(
				constant.MemoryIdentifier,
				mcp.Required(),
				mcp.Description("Memory ID"),
			),
			mcp.WithNumber(
				constant.ParentIdentifier,
				mcp.Description("Parent memory ID (omit to unset)"),
			),
		),
		s.setParent,
	)
}
