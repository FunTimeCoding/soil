package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMoveSymbols() {
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
				mcp.Description("Full import path of the destination package."),
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
			mcp.WithBoolean(
				"qualify_back_references",
				mcp.Description(
					"When a moved declaration references symbols staying in the source package, qualify those references and import the source package instead of refusing. Unexported back-reference targets still refuse - export them first.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.moveSymbols),
	)
}
