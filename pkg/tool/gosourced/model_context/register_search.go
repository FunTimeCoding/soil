package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSearch() {
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
			constant.FindLiterals,
			mcp.WithDescription(
				"Census of every composite literal and new() of a struct type across the module, before giving it a constructor. Sites outside the declaring package come back grouped by shape - the fields each sets in declaration order, & for an address, markers unkeyed, elided (type left out in a slice or map literal), value (no &), nested (not the sole value of an assignment or var) - with an exemplar and locations per group. Sites inside the declaring package and on the expected side of an assert are only counted. Read-only, no files change.",
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
				mcp.Description("Struct type name, e.g. Shape."),
			),
		),
		mcp.NewTypedToolHandler(s.findLiterals),
	)
}
