package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMoveMethods() {
	s.server.AddTool(
		mcp.NewTool(
			constant.MoveMethods,
			mcp.WithDescription(
				"Move methods of one type onto an existing type in another package, retyping their receivers. Every receiver field and sibling method a moved body uses must exist on the target with the same type, or move in the same batch. Callers outside the batch are routed through the source type's single field of the target type (x.M() becomes x.F.M()); without one the move refuses and names the call. Refuses a name the target already has and a method that helps the source type satisfy an interface. All-or-nothing; dry-run first.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the type.",
				),
			),
			mcp.WithString(
				constant.ParameterType,
				mcp.Required(),
				mcp.Description("Type whose methods move, e.g. Service."),
			),
			mcp.WithArray(
				"methods",
				mcp.Required(),
				mcp.Description(
					"Method names to move. Methods calling each other on the receiver move together.",
				),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description("Full import path of the destination package."),
			),
			mcp.WithString(
				"target_type",
				mcp.Required(),
				mcp.Description(
					"Existing type in the destination package that receives the methods.",
				),
			),
			mcp.WithString(
				"target_file",
				mcp.Description(
					"Destination file name inside the target package. Defaults to each method's source file name.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.moveMethods),
	)
}
