package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerUsers() {
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
}
