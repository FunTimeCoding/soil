package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerParameters() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveParameters,
			mcp.WithDescription(
				"Remove parameters from Go functions and methods, dropping the matching argument at every call site in the module. Batch: a parameter only passed on to another removed parameter in the batch counts as unused. Refuses a parameter still used in its body, a function used as a value or satisfying an interface, and any drop that would remove a call. Locals left unused are removed; leftover parameters and write-only fields are reported. All-or-nothing.",
			),
			mcp.WithArray(
				"functions",
				mcp.Required(),
				mcp.Description(
					"Functions and the parameters to remove from each.",
				),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"package_path": map[string]any{
								"type":        "string",
								"description": "Full import path of the package.",
							},
							"name": map[string]any{
								"type":        "string",
								"description": "Function or method name.",
							},
							"receiver": map[string]any{
								"type":        "string",
								"description": "Receiver type name for methods, e.g. Store.",
							},
							"parameters": map[string]any{
								"type":        "array",
								"items":       map[string]any{"type": "string"},
								"description": "Parameter names to remove.",
							},
						},
						"required": []string{
							"package_path",
							"name",
							"parameters",
						},
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
		mcp.NewTypedToolHandler(s.removeParameters),
	)
}
