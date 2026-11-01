package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerImports() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddImport,
			mcp.WithDescription("Add an import to a Go file."),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description("File path relative to module root."),
			),
			mcp.WithString(
				"import_path",
				mcp.Required(),
				mcp.Description("Import path to add, e.g. fmt."),
			),
			mcp.WithString("alias", mcp.Description("Optional import alias.")),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.addImport),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveImport,
			mcp.WithDescription("Remove an import from a Go file."),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description("File path relative to module root."),
			),
			mcp.WithString(
				"import_path",
				mcp.Required(),
				mcp.Description("Import path to remove."),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.removeImport),
	)
}
