package model_context

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
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
			mcp.WithDescription("Set the active module for this session."),
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
			constant.FindReferences,
			mcp.WithDescription(
				"Find all type-checked references to a Go symbol across the module, or census a file: per top-level symbol, the references outside the declaring file. Read-only, no files change. Counts are module-scoped - references from other modules are invisible, so a zero count only proves the symbol unused within this module.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the symbol or file.",
				),
			),
			mcp.WithString(
				"symbol",
				mcp.Description(
					"Symbol name to look up. Pass this or file, not both.",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description("Receiver type name for methods, e.g. Store."),
			),
			mcp.WithString(
				"file",
				mcp.Description(
					"File path relative to module root - censuses every top-level symbol declared in it. Pass this or symbol, not both.",
				),
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description(
					"Maximum locations returned per symbol. Defaults to 25; the exact total is always reported.",
				),
			),
			mcp.WithNumber(
				generative.ParameterOffset,
				mcp.Description(
					"Locations to skip before the limit window - page through large reference lists.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.findReferences),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MatchPattern,
			mcp.WithDescription(
				"Pre-flight check for a bulk edit: test every reference to a symbol against an expected shape, before editing anything. The pattern is plain Go - a function whose parameters declare the holes; block bodies in the pattern are wildcards. Sites whose enclosing statement matches count as matched; the rest come back grouped by normalized shape (literals as INT/STRING, other identifiers as IDENT) with a verbatim exemplar and locations per group - the sites that need hands instead of the mechanical edit. Read-only, no files change.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the symbol. Imported packages work too, e.g. fmt.",
				),
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description(
					"Symbol name whose references are tested. A type name with no receiver anchors every method on the type; the pattern then spells the anchor call with the type name standing for any method.",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description("Receiver type name for methods, e.g. Client."),
			),
			mcp.WithString(
				"pattern",
				mcp.Required(),
				mcp.Description(
					"Expected shape as a Go function with exactly one body statement, e.g. func pattern(c *client.Client, key string) { fmt.Println(c.DeleteComment(key)) }. Parameters are holes matching any expression assignable to their declared type; a single []any hole spread as c.Method(x...) matches any argument list.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.matchPattern),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ApplyPattern,
			mcp.WithDescription(
				"Apply a bulk edit to every site matching a pattern: run the match_pattern check, then rewrite each matched site's statement from the pattern shape to the replacement shape, holes carrying the site's own expressions. Imports are managed automatically. By default all-or-nothing: any unmatched or refused site means nothing is written and the report explains why; pass partial to rewrite the matched sites and take the remainder as hand work. Sites with comments in the statement refuse rather than lose them. Run with dry_run first.",
			),
			mcp.WithString(
				"package_path",
				mcp.Required(),
				mcp.Description(
					"Full import path of the package holding the symbol. Imported packages work too, e.g. fmt.",
				),
			),
			mcp.WithString(
				"symbol",
				mcp.Required(),
				mcp.Description(
					"Anchor symbol name. A type name with no receiver anchors every method on the type.",
				),
			),
			mcp.WithString(
				"receiver",
				mcp.Description("Receiver type name for methods, e.g. Client."),
			),
			mcp.WithString(
				"pattern",
				mcp.Required(),
				mcp.Description(
					"Current shape, as in match_pattern. May carry an import block when the replacement introduces packages the tool cannot resolve unambiguously.",
				),
			),
			mcp.WithString(
				"replacement",
				mcp.Required(),
				mcp.Description(
					"Target shape - a Go function with the same hole names and exactly one body statement, e.g. func replacement(c *client.Client, key string) { console.Emit(c.DeleteComment(key)) }.",
				),
			),
			mcp.WithBoolean(
				"partial",
				mcp.Description(
					"Rewrite the matched sites even when unmatched or refused sites remain. Off by default - any imperfection refuses the whole application.",
				),
			),
			mcp.WithBoolean(
				"dry_run",
				mcp.Description(
					"Report what the call would change without writing anything. Emits the same lines a real run does.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.applyPattern),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCalls,
			mcp.WithDescription(
				"Inventory of what a region of the module calls: every function and method referenced by packages under an import-path prefix, excluding calls within the region itself, grouped by callee with counts, largest first. The discovery step before a bulk edit - surfaces sibling names nobody thought to ask about (fmt.Print beside fmt.Println) and shows a region's blast radius. Read-only, no files change.",
			),
			mcp.WithString(
				"region",
				mcp.Required(),
				mcp.Description(
					"Import path prefix selecting the region, e.g. github.com/funtimecoding/soil/pkg/tool/gonetboxd.",
				),
			),
			mcp.WithNumber(
				generative.ParameterLimit,
				mcp.Description(
					"Maximum callees returned. Defaults to 100; the exact total is always reported.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.listCalls),
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
				mcp.Description("Receiver type name for methods, e.g. Store."),
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
	s.server.AddTool(
		mcp.NewTool(
			constant.RenameSymbol,
			mcp.WithDescription(
				"Rename a Go function, method, type, or constant. Updates all references across the module.",
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
				mcp.Description("Receiver type name for methods, e.g. Store."),
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
				"type",
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
				mcp.Description("Function or method name to extract."),
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
