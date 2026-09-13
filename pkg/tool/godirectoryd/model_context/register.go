package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListUser,
			mcp.WithDescription(
				"List every user in the directory with account, name, mail and stable unique identifier.",
			),
		),
		mcp.NewTypedToolHandler(s.listUser),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateUser,
			mcp.WithDescription(
				"Create a directory user. Account is the login name; surname is required by the person schema.",
			),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name, unique in the directory"),
			),
			mcp.WithString(
				constant.NameField,
				mcp.Required(),
				mcp.Description("Display name"),
			),
			mcp.WithString(
				constant.SurnameField,
				mcp.Required(),
				mcp.Description("Family name"),
			),
			mcp.WithString(constant.MailField, mcp.Description("Mail address")),
			mcp.WithString(
				constant.PasswordField,
				mcp.Description("Initial password; omit to create without one"),
			),
		),
		mcp.NewTypedToolHandler(s.createUser),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ModifyUser,
			mcp.WithDescription(
				"Change a directory user. Omitted fields are left untouched.",
			),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name"),
			),
			mcp.WithString(
				constant.NameField,
				mcp.Description("New display name"),
			),
			mcp.WithString(
				constant.SurnameField,
				mcp.Description("New family name"),
			),
			mcp.WithString(
				constant.MailField,
				mcp.Description("New mail address"),
			),
		),
		mcp.NewTypedToolHandler(s.modifyUser),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteUser,
			mcp.WithDescription("Remove a directory user permanently."),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name"),
			),
		),
		mcp.NewTypedToolHandler(s.deleteUser),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetPassword,
			mcp.WithDescription(
				"Set a directory user's password through the password modify operation.",
			),
			mcp.WithString(
				constant.AccountField,
				mcp.Required(),
				mcp.Description("Login name"),
			),
			mcp.WithString(
				constant.PasswordField,
				mcp.Required(),
				mcp.Description("New password"),
			),
		),
		mcp.NewTypedToolHandler(s.setPassword),
	)
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
