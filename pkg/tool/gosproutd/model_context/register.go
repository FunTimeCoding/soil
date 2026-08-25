package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSeeds,
			mcp.WithDescription(
				"List all seeds sorted by priority. Each seed carries its file modification time.",
			),
		),
		mcp.NewTypedToolHandler(s.listSeeds),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReorderSeed,
			mcp.WithDescription(
				"Move a seed to a specific position in the priority list.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Seed name"),
			),
			mcp.WithNumber(
				"position",
				mcp.Required(),
				mcp.Description("Target position (1 = top)"),
			),
		),
		mcp.NewTypedToolHandler(s.reorderSeed),
	)
}
