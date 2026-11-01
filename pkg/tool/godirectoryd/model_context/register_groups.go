package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerGroups() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListGroup,
			mcp.WithDescription(
				"List every group in the directory with its number and members.",
			),
		),
		mcp.NewTypedToolHandler(s.listGroup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateGroup,
			mcp.WithDescription(
				"Create a group. The group number is allocated automatically.",
			),
			mcp.WithString(
				constant.NameField,
				mcp.Required(),
				mcp.Description("Group name"),
			),
		),
		mcp.NewTypedToolHandler(s.createGroup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteGroup,
			mcp.WithDescription("Remove a group permanently."),
			mcp.WithString(
				constant.NameField,
				mcp.Required(),
				mcp.Description("Group name"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteGroup),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddMember,
			mcp.WithDescription(
				"Add an existing user to a group. Refuses when the user is not in the directory.",
			),
			mcp.WithString(
				constant.NameField,
				mcp.Required(),
				mcp.Description("Group name"),
			),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name of the user to add"),
			),
		),
		mcp.NewTypedToolHandler(s.addMember),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveMember,
			mcp.WithDescription("Remove a user from a group."),
			mcp.WithString(
				constant.NameField,
				mcp.Required(),
				mcp.Description("Group name"),
			),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name of the user to remove"),
			),
		),
		mcp.NewTypedToolHandler(s.removeMember),
	)
}
