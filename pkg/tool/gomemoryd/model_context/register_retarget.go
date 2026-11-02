package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRetarget() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RetargetReferences,
			mcp.WithDescription(constant.RetargetSummary),
			mcp.WithArray(
				"paths",
				mcp.Description(
					"Moved files or directories, absolute or relative to the reference root (e.g. doc/ai/seed/old.md to doc/ai/seed/new.md, ../soil/pkg/alfa to ../soil/pkg/bravo).",
				),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"from": map[string]any{
								"type":        "string",
								"description": "Path before the move.",
							},
							"to": map[string]any{
								"type":        "string",
								"description": "Path after the move.",
							},
						},
						"required": []string{"from", "to"},
					},
				),
			),
			mcp.WithArray(
				"symbols",
				mcp.Description(
					"Moved or renamed Go symbols, by package directory as go: references spell it.",
				),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"directory": map[string]any{
								"type":        "string",
								"description": "Package directory before, e.g. ../soil/pkg/alfa.",
							},
							"name": map[string]any{
								"type":        "string",
								"description": "Name before, Type.Method for a method.",
							},
							"new_directory": map[string]any{
								"type":        "string",
								"description": "Package directory after.",
							},
							"new_name": map[string]any{
								"type":        "string",
								"description": "Name after.",
							},
						},
						"required": []string{
							"directory",
							"name",
							"new_directory",
							"new_name",
						},
					},
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what would change without writing anything.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.retargetReferences),
	)
}
