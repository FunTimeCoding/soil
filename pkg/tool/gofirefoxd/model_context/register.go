package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTabs,
			mcp.WithDescription(
				"List all open Firefox tabs with URL, title, and active state",
			),
		),
		mcp.NewTypedToolHandler(s.ListTabs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadTab,
			mcp.WithDescription(
				"Read the text content of a Firefox tab. Extracts readable article text by default (strips navigation, ads, footers). Falls back to raw innerText if readability extraction fails. Pass raw=true to skip readability and get innerText directly - better for dashboards, web apps, and non-article pages.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithBoolean(
				"raw",
				mcp.Description(
					"Skip readability, return raw innerText (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.ReadTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTab,
			mcp.WithDescription("Open a URL in a new Firefox tab"),
			mcp.WithString(
				"url",
				mcp.Description("URL to open (defaults to about:blank)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CloseTab,
			mcp.WithDescription("Close a Firefox tab"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.CloseTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Navigate,
			mcp.WithDescription("Navigate a Firefox tab to a URL"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to navigate to"),
			),
		),
		mcp.NewTypedToolHandler(s.Navigate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NavigateBack,
			mcp.WithDescription("Go back in a Firefox tab's history"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.NavigateBack),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GroupTabs,
			mcp.WithDescription(
				"Group Firefox tabs together. Creates a new group or adds to an existing group. Optionally set title and color.",
			),
			mcp.WithArray(
				"tab_ids",
				mcp.Required(),
				mcp.Description("Array of tab identifiers to group"),
			),
			mcp.WithNumber(
				"group_id",
				mcp.Description(
					"Existing group ID to add tabs to (omit to create new group)",
				),
			),
			mcp.WithString("title", mcp.Description("Group title")),
			mcp.WithString(
				"color",
				mcp.Description(
					"Group color (grey, blue, red, yellow, green, pink, purple, cyan, orange)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.GroupTabs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UngroupTab,
			mcp.WithDescription("Remove a tab from its group"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.UngroupTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateGroup,
			mcp.WithDescription("Update a tab group's title or color"),
			mcp.WithNumber(
				"group_id",
				mcp.Required(),
				mcp.Description("Group identifier"),
			),
			mcp.WithString("title", mcp.Description("New group title")),
			mcp.WithString(
				"color",
				mcp.Description(
					"New group color (grey, blue, red, yellow, green, pink, purple, cyan, orange)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateGroup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.WaitFor,
			mcp.WithDescription(
				"Wait for a Firefox tab to finish loading, and optionally until a CSS selector appears. Essential over Tor, where a page takes seconds to settle and a read issued too early returns the previous page. Fails if the selector never appears.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"selector",
				mcp.Description(
					"CSS selector to wait for (omit to wait only for page load)",
				),
			),
			mcp.WithNumber(
				"timeout",
				mcp.Description("Timeout in milliseconds (default 30000)"),
			),
		),
		mcp.NewTypedToolHandler(s.WaitFor),
	)
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
	s.server.AddTool(
		mcp.NewTool(
			constant.Screenshot,
			mcp.WithDescription(
				"Capture a PNG of a Firefox tab and write it to the download directory, returning the path to read. Use it to see what text extraction cannot: CAPTCHA and Cloudflare interstitials, which are frequent over Tor and look like an empty page otherwise. Firefox can only capture the visible tab, so this activates the target tab first.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.Screenshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadFile,
			mcp.WithDescription(
				"Download a URL through the browser, carrying its cookies and its proxy. Always prefer this over fetching a URL outside the browser, which would bypass Tor. Saves without a prompt and uniquifies a colliding name.",
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to download"),
			),
			mcp.WithString(
				"filename",
				mcp.Description(
					"Target name, relative to the browser's download directory",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadFile),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadStatus,
			mcp.WithDescription(
				"Report Firefox downloads with their state, bytes received and final path. Omit download_id for the ten most recent.",
			),
			mcp.WithNumber(
				"download_id",
				mcp.Description(
					"Download identifier (omit for the most recent downloads)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadStatus),
	)
}
