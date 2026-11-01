package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerReferences() {
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
				mcp.Description(
					"Type name for methods and struct fields, e.g. Store.",
				),
			),
			mcp.WithString(
				"file",
				mcp.Description(
					"File path relative to module root - censuses every top-level symbol declared in it. Pass this or symbol, not both.",
				),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Description(
					"Maximum locations returned per symbol. Defaults to 25; the exact total is always reported.",
				),
			),
			mcp.WithNumber(
				generativeConstant.ParameterOffset,
				mcp.Description(
					"Locations to skip before the limit window - page through large reference lists.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.findReferences),
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
				generativeConstant.ParameterLimit,
				mcp.Description(
					"Maximum callees returned. Defaults to 100; the exact total is always reported.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.listCalls),
	)
}
