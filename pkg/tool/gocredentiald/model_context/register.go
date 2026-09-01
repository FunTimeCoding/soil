package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListEntries,
			mcp.WithDescription(
				"List every entry: identifier, group path, title, user, locator, modification time. Passwords never appear here.",
			),
		),
		mcp.NewTypedToolHandler(s.listEntries),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetEntry,
			mcp.WithDescription(
				"One entry's structure. Only username and locator values arrive readable - passwords, notes, and custom field values are masked; use reveal_password when a password is genuinely needed.",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Entry identifier from list or search"),
			),
		),
		mcp.NewTypedToolHandler(s.getEntry),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RevealPassword,
			mcp.WithDescription(
				"Reveal one entry's password. Deliberate, single-entry, on demand only.",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Entry identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.revealPassword),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchEntries,
			mcp.WithDescription(
				"Case-insensitive search across title, user, locator, notes, and group path.",
			),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("Search text"),
			),
		),
		mcp.NewTypedToolHandler(s.searchEntries),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AuditEntries,
			mcp.WithDescription(
				"Staleness report: entries unmodified beyond the threshold, empty users, empty passwords, duplicate title+user pairs. Without bucket: per-bucket totals plus a small sample. With bucket: paged rows from that bucket.",
			),
			mcp.WithNumber(
				"stale_years",
				mcp.Description("Staleness threshold in years (default 3)"),
			),
			mcp.WithString(
				"bucket",
				mcp.Description(
					"Bucket to page: stale, empty_user, empty_password, duplicates",
				),
			),
			mcp.WithNumber("limit", mcp.Description("Page size (default 100)")),
			mcp.WithNumber("offset", mcp.Description("Page offset")),
		),
		mcp.NewTypedToolHandler(s.auditEntries),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateEntry,
			mcp.WithDescription(
				"Create an entry in a group. Fields are KEY=VALUE strings; Password becomes a protected field.",
			),
			mcp.WithString(
				"group",
				mcp.Required(),
				mcp.Description("Group path, e.g. Root/Web"),
			),
			mcp.WithString(
				"title",
				mcp.Required(),
				mcp.Description("Entry title"),
			),
			mcp.WithArray(
				"fields",
				mcp.Description("KEY=VALUE field strings"),
				mcp.WithStringItems(),
			),
		),
		mcp.NewTypedToolHandler(s.createEntry),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateEntry,
			mcp.WithDescription(
				"Set fields on an entry. Fields are KEY=VALUE strings; existing keys update, new keys append.",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Entry identifier"),
			),
			mcp.WithArray(
				"fields",
				mcp.Required(),
				mcp.Description("KEY=VALUE field strings"),
				mcp.WithStringItems(),
			),
		),
		mcp.NewTypedToolHandler(s.updateEntry),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MoveEntry,
			mcp.WithDescription("Move an entry to another group."),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Entry identifier"),
			),
			mcp.WithString(
				"group",
				mcp.Required(),
				mcp.Description("Target group path"),
			),
		),
		mcp.NewTypedToolHandler(s.moveEntry),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteEntry,
			mcp.WithDescription("Delete an entry permanently."),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Entry identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteEntry),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.LoadGroup,
			mcp.WithDescription(
				"Load an environment variable group by name from the Environment group: custom fields returned as NAME=value pairs, revealed.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Environment group entry title"),
			),
		),
		mcp.NewTypedToolHandler(s.loadGroup),
	)
}
