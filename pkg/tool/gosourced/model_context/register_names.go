package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerNames() {
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameSymbol,
			mcp.WithDescription(
				"Rename a Go function, method, struct field, type, or constant. Updates all references across the module, composite literal keys included. Embedded fields refuse - rename the type instead.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description("Full import path of the package."),
			),
			mcp.WithString(
				"old_name",
				mcp.Required(),
				mcp.Description("Current name of the symbol."),
			),
			mcp.WithString(
				"new_name",
				mcp.Required(),
				mcp.Description("New name for the symbol."),
			),
			mcp.WithString(
				"receiver",
				mcp.Description(
					"Type name for methods and struct fields, e.g. Store.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.renameSymbol),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ChangeVisibility,
			mcp.WithDescription(
				"Change the visibility of a Go function, method, struct field, type, or constant by toggling its first letter case. Updates all references across the module.",
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description("Symbol name, e.g. IsGeneratedHeader."),
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
					"Type name for methods and struct fields, e.g. Store.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.changeVisibility),
	)
}
