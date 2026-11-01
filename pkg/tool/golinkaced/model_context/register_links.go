package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerLinks() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLinks,
			mcp.WithDescription("List links. Optionally filter by list ID."),
			mcp.WithNumber(
				constant.ListParameter,
				mcp.Description("List ID to filter by."),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListLinks),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddLink,
			mcp.WithDescription("Add a bookmark. List and tags accept names."),
			mcp.WithString(
				"link",
				mcp.Required(),
				mcp.Description("URL to bookmark."),
			),
			mcp.WithString(
				"name",
				mcp.Description("Display name for the bookmark."),
			),
			mcp.WithString(
				constant.ListParameter,
				mcp.Description("List name or ID."),
			),
			mcp.WithArray(
				"tags",
				mcp.WithStringItems(),
				mcp.Description("Tag names to apply."),
			),
		),
		mcp.NewTypedToolHandler(s.AddLink),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.EditLink,
			mcp.WithDescription(
				fmt.Sprintf(
					"Edit a bookmark. Supports replace and incremental tag/list changes. %s %s",
					"tags/lists replace all; add_tags/remove_tags and add_lists/remove_lists modify incrementally.",
					"Replace wins over incremental if both provided.",
				),
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Link ID to edit."),
			),
			mcp.WithString("name", mcp.Description("New display name.")),
			mcp.WithString("link", mcp.Description("New URL.")),
			mcp.WithString("description", mcp.Description("New description.")),
			mcp.WithArray(
				"tags",
				mcp.WithStringItems(),
				mcp.Description("Replace all tags with these names."),
			),
			mcp.WithArray(
				"add_tags",
				mcp.WithStringItems(),
				mcp.Description("Tag names to add."),
			),
			mcp.WithArray(
				"remove_tags",
				mcp.WithStringItems(),
				mcp.Description("Tag names to remove."),
			),
			mcp.WithArray(
				"lists",
				mcp.WithStringItems(),
				mcp.Description("Replace all lists with these names."),
			),
			mcp.WithArray(
				"add_lists",
				mcp.WithStringItems(),
				mcp.Description("List names to add."),
			),
			mcp.WithArray(
				"remove_lists",
				mcp.WithStringItems(),
				mcp.Description("List names to remove."),
			),
		),
		mcp.NewTypedToolHandler(s.EditLink),
	)
}
