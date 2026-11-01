package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerActions() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Evaluate,
			mcp.WithDescription(
				"Evaluate a JavaScript expression in a Firefox tab and return its value as JSON. The way to read links, since read_tab discards every href: `Array.from(document.querySelectorAll('a')).map(a => ({text: a.innerText, href: a.href}))` returns a compact result instead of a page-sized blob. For statements, use an arrow function: `(() => { ...; return x })()`. The value must be JSON-serializable - returning a DOM node fails.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"expression",
				mcp.Required(),
				mcp.Description("JavaScript expression to evaluate"),
			),
		),
		mcp.NewTypedToolHandler(s.Evaluate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Click,
			mcp.WithDescription(
				"Click the element matching a CSS selector in a Firefox tab. Scrolls it into view first. Returns the matched element's tag and text so the call confirms what it hit, and fails if the selector matches nothing.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"selector",
				mcp.Required(),
				mcp.Description("CSS selector of the element to click"),
			),
		),
		mcp.NewTypedToolHandler(s.Click),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Fill,
			mcp.WithDescription(
				"Type a value into the element matching a CSS selector in a Firefox tab. Simulates real typing by default (focus, select all, insert text), which is what forms that only react to genuine input events need. Pass direct=true to set the value outright and dispatch input and change events instead - faster, but some forms ignore it.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"selector",
				mcp.Required(),
				mcp.Description("CSS selector of the field to fill"),
			),
			mcp.WithString("value", mcp.Description("Value to type")),
			mcp.WithBoolean(
				constant.Direct,
				mcp.Description(
					"Set the value directly instead of simulating typing (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Fill),
	)
}
