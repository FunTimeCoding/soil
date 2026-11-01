package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerModules() {
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
			mcp.WithDescription("Set the active module for this session."),
			mcp.WithString(
				constant.ParameterModule,
				mcp.Required(),
				mcp.Description("Module name from list_modules."),
			),
		),
		mcp.NewTypedToolHandler(s.useModule),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameModule,
			mcp.WithDescription(
				"Rewrite every import of a Go module to a new module path - the source half of a major version bump. Checks the new module first: every symbol this module actually uses is looked up at the target and its type compared, so removed symbols and changed signatures are reported before anything is written. Refuses on breakage unless force is set. Does not touch go.mod - run go get <new_module_path>@<version> before the call so the target resolves, and go mod tidy after.",
			),
			mcp.WithString(
				"module_path",
				mcp.Required(),
				mcp.Description(
					"Current module path, e.g. gitlab.com/gitlab-org/api/client-go/v2.",
				),
			),
			mcp.WithString(
				"new_module_path",
				mcp.Required(),
				mcp.Description(
					"New module path, e.g. gitlab.com/gitlab-org/api/client-go/v3. Subpackage paths follow the prefix.",
				),
			),
			mcp.WithBoolean(
				"force",
				mcp.Description(
					"Rewrite even when symbols are missing or signatures changed. The breakage list is still reported - it becomes the hand-work list.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does, breakages included.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.renameModule),
	)
}
