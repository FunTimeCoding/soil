package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPackages() {
	s.server.AddTool(
		mcp.NewTool(
			constant.MovePackage,
			mcp.WithDescription(
				"Move a package directory to another location in the module. Rewrites every import of the package and its subpackages across the module, then moves the directory with everything in it. The package name stays - the last path segment must match.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description("Full import path of the package to move."),
			),
			mcp.WithString(
				"target_package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the destination. Last segment must match the source.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.movePackage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RenamePackage,
			mcp.WithDescription(
				"Rename a package: the package clause in every file (including test variants), the directory, all import paths of it and its subpackages, and every unaliased qualifier across the module. Aliased imports keep their alias.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description("Full import path of the package to rename."),
			),
			mcp.WithString(
				"new_name",
				mcp.Required(),
				mcp.Description(
					"New package name, e.g. depot. Becomes the new last path segment.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.renamePackage),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RenamePackageClause,
			mcp.WithDescription(
				"Rename only the package clause: the package line in every file (including test variants) and every unaliased qualifier across the module. The directory and import paths stay unchanged. Aliased imports keep their alias.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description("Full import path of the package to rename."),
			),
			mcp.WithString(
				"new_name",
				mcp.Required(),
				mcp.Description(
					"New package name. The directory and import path keep their names.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.renamePackageClause),
	)
}
