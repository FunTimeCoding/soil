package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerWrite() {
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
}
