package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerConstructors() {
	s.server.AddTool(
		mcp.NewTool(
			constant.IntroduceConstructor,
			mcp.WithDescription(
				"Give a struct a constructor and rewrite its literal sites outside the declaring package to call it. The name follows the package: New in new.go when the struct is alone, NewFoo in new_foo.go in a bag of data structs; a struct sharing its package with a receiver struct refuses with the extract_type move to make first. Parameters are field names in declaration order - none by default (blank constructor, add parameters as needed), a list, or the fields every site sets. Only sites setting exactly the parameters are rewritten, an expression swap that is safe anywhere; every other site is reported with its reason (value literal, misses a field, sets fields beyond the constructor, evaluation order of calls, comments inside). Run find_literals first and dry_run before writing.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package declaring the type.",
				),
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Exported struct type name, e.g. Shape."),
			),
			mcp.WithArray(
				"parameters",
				mcp.Description(
					"Field names to take as parameters; they appear in declaration order. Excludes parameters_from_sites.",
				),
				mcp.Items(map[string]any{"type": "string"}),
			),
			mcp.WithBoolean(
				"parameters_from_sites",
				mcp.Description(
					"Take as parameters the fields every outside site sets.",
				),
			),
			mcp.WithBoolean(
				"assign_rest",
				mcp.Description(
					"Also rewrite sites setting more than the parameters when they are the sole value of an assignment or var: the constructor call, then one field assignment per remaining field.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.introduceConstructor),
	)
}
