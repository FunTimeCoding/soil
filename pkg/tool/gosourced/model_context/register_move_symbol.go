package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMoveSymbol() {
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
				mcp.Description("Symbol name to move, e.g. itemFields."),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description("Full import path of the destination package."),
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
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.moveSymbol),
	)
}
