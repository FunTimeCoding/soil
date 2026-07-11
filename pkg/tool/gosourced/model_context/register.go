package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListModules,
			mcp.WithDescription("List available modules."),
		),
		mcp.NewTypedToolHandler(s.listModules),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UseModule,
			mcp.WithDescription(
				"Set the active module for this session.",
			),
			mcp.WithString(
				"module",
				mcp.Required(),
				mcp.Description("Module name from list_modules."),
			),
		),
		mcp.NewTypedToolHandler(s.useModule),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ChangeVisibility,
			mcp.WithDescription(
				"Change the visibility of a Go function, method, type, or constant by toggling its first letter case. Updates all references across the module.",
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description(
					"Symbol name, e.g. IsGeneratedHeader.",
				),
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package, e.g. github.com/funtimecoding/soil/pkg/lint.",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description(
					"Receiver type name for methods, e.g. Store.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.changeVisibility),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameSymbol,
			mcp.WithDescription(
				"Rename a Go function, method, type, or constant. Updates all references across the module.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package.",
				),
			),
			mcp.WithString(
				"old_name",
				mcp.Required(),
				mcp.Description(
					"Current name of the symbol.",
				),
			),
			mcp.WithString(
				"new_name",
				mcp.Required(),
				mcp.Description(
					"New name for the symbol.",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description(
					"Receiver type name for methods, e.g. Store.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.renameSymbol),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MoveSymbol,
			mcp.WithDescription(
				"Move a top-level Go constant, variable, type, or function to another package. Exports the symbol if needed, qualifies all references, and manages imports on both sides.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the symbol.",
				),
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description(
					"Symbol name to move, e.g. itemFields.",
				),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the destination package.",
				),
			),
			mcp.WithString(
				"target_file",
				mcp.Description(
					"Destination file name inside the target package, e.g. constant.go. Appends when the file exists, creates it otherwise. Defaults to the snake_case symbol name.",
				),
			),
			mcp.WithBoolean(
				"create",
				mcp.Description(
					"Create the destination package when it does not exist. Refuses by default.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.moveSymbol),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MoveSymbols,
			mcp.WithDescription(
				"Move multiple top-level Go symbols to another package in one call. Symbols moving together may reference each other - those references stay unqualified. Constants merge into one declaration group. All-or-nothing: any refused symbol blocks the whole batch.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the symbols.",
				),
			),
			mcp.WithArray(
				"symbols",
				mcp.Description(
					"Symbol names to move. Pass this or file, not both.",
				),
			),
			mcp.WithString(
				"file",
				mcp.Description(
					"File path relative to module root - moves every top-level symbol declared in it. Pass this or symbols, not both.",
				),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the destination package.",
				),
			),
			mcp.WithString(
				"target_file",
				mcp.Description(
					"Destination file name inside the target package, e.g. constant.go. Appends when the file exists, creates it otherwise. Defaults to the snake_case symbol name per symbol.",
				),
			),
			mcp.WithBoolean(
				"create",
				mcp.Description(
					"Create the destination package when it does not exist. Refuses by default.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.moveSymbols),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ExtractToFile,
			mcp.WithDescription(
				"Extract a function or method from a file into its own file. Carries needed imports, removes unused imports from the source.",
			),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description(
					"File path relative to module root, e.g. pkg/tool/gosourced/service/change_visibility.go.",
				),
			),
			mcp.WithString(
				"function",
				mcp.Required(),
				mcp.Description(
					"Function or method name to extract.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.extractToFile),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddImport,
			mcp.WithDescription(
				"Add an import to a Go file.",
			),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description(
					"File path relative to module root.",
				),
			),
			mcp.WithString(
				"import_path",
				mcp.Required(),
				mcp.Description(
					"Import path to add, e.g. fmt.",
				),
			),
			mcp.WithString(
				"alias",
				mcp.Description(
					"Optional import alias.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.addImport),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveImport,
			mcp.WithDescription(
				"Remove an import from a Go file.",
			),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description(
					"File path relative to module root.",
				),
			),
			mcp.WithString(
				"import_path",
				mcp.Required(),
				mcp.Description(
					"Import path to remove.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.removeImport),
	)
}
