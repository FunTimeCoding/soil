package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerGraph() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RelateMemories,
			mcp.WithDescription(
				"Create a relation between two memories. Relations are bidirectional - get_memory on either shows the other under related. Relating an existing pair with a type retypes the edge.",
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
			mcp.WithString(
				constant.Type,
				mcp.Description(
					"Relation type: affinity, informs, grounds, mechanism, sequence, deep-dive (optional; affinity and informs read undirected, the rest source to target)",
				),
			),
		),
		s.relate,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UnrelateMemories,
			mcp.WithDescription(
				"Remove a relation between two memories. The stored row is directional - pass the ids in the order the relation was created; the error message says when to try the reverse. Permanent, no history.",
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
		s.unrelate,
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
