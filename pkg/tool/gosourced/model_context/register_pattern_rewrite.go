package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPatternRewrite() {
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
					"Target shape - a Go function with the same hole names and exactly one body statement, e.g. func replacement(c *client.Client, t *terminal.Terminal, key string) { t.Emit(c.DeleteComment(key)) }.",
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
}
