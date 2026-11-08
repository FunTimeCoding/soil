package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRenameReceivers() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameReceivers,
			mcp.WithDescription(
				"Rename the receiver of every method of a type, and each use of it, to one name. Uses are found by type-checked identity, never by spelling. An unnamed or blank receiver stays unnamed. Refuses a name that is not an identifier or is built in, a type without methods, and any method where the name is already taken - a local, a parameter, or a package-level object or import that the receiver would otherwise shadow. All-or-nothing across the batch.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package declaring the types.",
				),
			),
			mcp.WithArray(
				"receivers",
				mcp.Required(),
				mcp.Description("Types and the receiver name each takes."),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"type": map[string]any{
								"type":        "string",
								"description": "Type whose methods' receivers are renamed, e.g. Service.",
							},
							"name": map[string]any{
								"type":        "string",
								"description": "New receiver name, e.g. s.",
							},
						},
						"required": []string{"type", "name"},
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
		mcp.NewTypedToolHandler(s.renameReceivers),
	)
}
