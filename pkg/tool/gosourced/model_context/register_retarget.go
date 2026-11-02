package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRetarget() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RetargetReferences,
			mcp.WithDescription(
				"Fix references after files or directories were moved by hand (git mv): rewrites markdown paths and go: spans in this module and every module requiring it, and the memories citing them, as gosourced's own moves do. Works after the move happened. Paths are relative to the module root.",
			),
			mcp.WithArray(
				"paths",
				mcp.Required(),
				mcp.Description("Moved files or directories."),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"from": map[string]any{
								"type":        "string",
								"description": "Path before the move, e.g. doc/ai/seed/old.md.",
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
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.retargetReferences),
	)
}
