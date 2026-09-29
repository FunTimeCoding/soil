package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
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
			constant.Search,
			mcp.WithDescription(
				"Search links, lists, and tags. Optionally filter by type.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search query string."),
			),
			mcp.WithString(
				"type",
				mcp.Description("Entity type: link, list, tag. Omit for all."),
			),
		),
		mcp.NewTypedToolHandler(s.Search),
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
	s.server.AddTool(
		mcp.NewTool(
			constant.Delete,
			mcp.WithDescription(
				"Delete a resource. Type: link, list, tag, note, or branch (deletes list and all its links).",
			),
			mcp.WithString(
				"type",
				mcp.Required(),
				mcp.Description("Resource type: link, list, tag, note, branch."),
			),
			mcp.WithNumber(
				"identifier",
				mcp.Required(),
				mcp.Description("Resource ID to delete."),
			),
		),
		mcp.NewTypedToolHandler(s.Delete),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLists,
			mcp.WithDescription("List all LinkAce lists."),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListLists),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription("List all LinkAce tags."),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results to return."),
			),
			mcp.WithNumber(
				"offset",
				mcp.Description("Number of results to skip."),
			),
		),
		mcp.NewTypedToolHandler(s.ListTags),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNotes,
			mcp.WithDescription("List notes on a link."),
			mcp.WithNumber(
				"link_identifier",
				mcp.Required(),
				mcp.Description("Link ID to list notes for."),
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
		mcp.NewTypedToolHandler(s.ListNotes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddNote,
			mcp.WithDescription("Add a note to a link."),
			mcp.WithNumber(
				"link_identifier",
				mcp.Required(),
				mcp.Description("Link ID to add the note to."),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description("Note text."),
			),
		),
		mcp.NewTypedToolHandler(s.AddNote),
	)
}
