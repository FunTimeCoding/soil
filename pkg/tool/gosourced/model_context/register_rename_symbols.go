package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRenameSymbols() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameSymbols,
			mcp.WithDescription(
				"Rename several Go symbols declared in one package with a single load. Every rename is checked before anything is written, as rename_symbol checks one. All-or-nothing: any refused rename blocks the whole batch, as does a symbol listed twice or two renames to the same name. A rename onto a name another rename in the batch frees still refuses - run those in order.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package declaring the symbols.",
				),
			),
			mcp.WithArray(
				"renames",
				mcp.Required(),
				mcp.Description("Renames to apply."),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"old_name": map[string]any{
								"type":        "string",
								"description": "Current name of the symbol.",
							},
							"new_name": map[string]any{
								"type":        "string",
								"description": "New name for the symbol.",
							},
							"receiver": map[string]any{
								"type":        "string",
								"description": "Type name for methods and struct fields, e.g. Store.",
							},
						},
						"required": []string{"old_name", "new_name"},
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
		mcp.NewTypedToolHandler(s.renameSymbols),
	)
}
