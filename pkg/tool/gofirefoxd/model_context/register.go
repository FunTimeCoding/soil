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
			mcp.WithObject(
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
			mcp.WithString(
				"title",
				mcp.Description("Group title"),
			),
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
			mcp.WithString(
				"title",
				mcp.Description("New group title"),
			),
			mcp.WithString(
				"color",
				mcp.Description(
					"New group color (grey, blue, red, yellow, green, pink, purple, cyan, orange)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateGroup),
	)
}
