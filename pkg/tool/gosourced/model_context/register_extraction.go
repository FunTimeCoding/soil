package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerExtraction() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ExtractType,
			mcp.WithDescription(
				"Move a named type and its method set to another package as one unit. Members referenced outside the moving set (methods, struct fields) are exported as needed and all references renamed. Declarations land in target files matching their source basenames unless target_file collapses them into one.",
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
				mcp.Description("Named type to extract, e.g. Store."),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description("Full import path of the destination package."),
			),
			mcp.WithString(
				"target_file",
				mcp.Description(
					"Destination file name inside the target package. Defaults to preserving each declaration's source basename.",
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
		mcp.NewTypedToolHandler(s.extractType),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ExtractToFile,
			mcp.WithDescription(
				"Extract a function, method, or type from a file into its own file. Carries needed imports, removes unused imports from the source. When one identity remains in the source, the source is renamed to match it.",
			),
			mcp.WithString(
				"file",
				mcp.Required(),
				mcp.Description(
					"File path relative to module root, e.g. pkg/tool/gosourced/service/change_visibility.go.",
				),
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description(
					"Function, method, or type name to extract. Types move without their methods - methods are their own identities. Grouped type declarations refuse.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.extractToFile),
	)
}
