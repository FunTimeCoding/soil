package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerGroups() {
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
}
