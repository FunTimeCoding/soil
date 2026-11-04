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
	s.server.AddTool(
		mcp.NewTool(
			constant.AddParameters,
			mcp.WithDescription(
				"Add a parameter to Go functions and methods at a position, passing an argument expression at every call site in the module, replacing modules included. Package qualifiers in the type and the argument resolve to the one package of that name in the build graph, or to a listed import; imports are added where needed. Each call site type-checks the argument in its own scope: a site where it does not resolve, or is not assignable to the parameter type, is refused by position. Refuses a function used as a value or satisfying an interface, a call passing a multi-value result, unnamed parameters, a position out of range or past a variadic parameter, and a name the body already uses. All-or-nothing.",
			),
			mcp.WithArray(
				"functions",
				mcp.Required(),
				mcp.Description("Functions and the parameter to add to each."),
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
							"parameter": map[string]any{
								"type":        "string",
								"description": "Name of the new parameter, e.g. x.",
							},
							"type": map[string]any{
								"type":        "string",
								"description": "Go type expression, e.g. context.Context.",
							},
							"position": map[string]any{
								"type":        "integer",
								"description": "Zero-based position among the parameters.",
							},
							"argument": map[string]any{
								"type":        "string",
								"description": "Go expression passed at every call site, e.g. x or context.Background().",
							},
						},
						"required": []string{
							"package_path",
							"name",
							"parameter",
							"type",
							"position",
							"argument",
						},
					},
				),
			),
			mcp.WithArray(
				"imports",
				mcp.Items(map[string]any{"type": "string"}),
				mcp.Description(
					"Import paths for package qualifiers that are ambiguous or not yet in the build graph.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.addParameters),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NameParameters,
			mcp.WithDescription(
				"Name the unnamed or blank parameter of a given type in every function declaration and literal in scope, e.g. _ context.Context becomes x context.Context. Named parameters are left alone; when the list was all unnamed, the other parameters become _. Refuses a function where the name is already declared in, used from or returned by its body or signature, and one with more than one such parameter. All-or-nothing.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package; a trailing /... covers its subpackages.",
				),
			),
			mcp.WithArray(
				"functions",
				mcp.Description(
					"Narrow the scope to these functions and the literals inside them.",
				),
				mcp.Items(
					map[string]any{
						"type": "object",
						"properties": map[string]any{
							"name": map[string]any{
								"type":        "string",
								"description": "Function or method name.",
							},
							"receiver": map[string]any{
								"type":        "string",
								"description": "Receiver type name for methods, e.g. Server.",
							},
						},
						"required": []string{"name"},
					},
				),
			),
			mcp.WithString(
				constant.ParameterType,
				mcp.Required(),
				mcp.Description("Go type expression, e.g. context.Context."),
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Name to give the parameter, e.g. x."),
			),
			mcp.WithArray(
				"imports",
				mcp.Items(map[string]any{"type": "string"}),
				mcp.Description(
					"Import paths for package qualifiers that are ambiguous or not yet in the build graph.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.nameParameters),
	)
}
