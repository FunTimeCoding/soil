package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func createUserForm() gomponents.Node {
	return html.Form(
		html.Method("post"),
		html.Action(constant.CreateUserPath),
		html.H2(gomponents.Text("Create user")),
		field(constant.AccountField, "Account", constant.TextInputType, true),
		field(constant.NameField, "Name", constant.TextInputType, true),
		field(constant.SurnameField, "Surname", constant.TextInputType, true),
		field(constant.MailField, "Mail", constant.MailInputType, false),
		field(
			constant.PasswordField,
			"Password",
			constant.PasswordInputType,
			false,
		),
		html.Button(html.Type("submit"), gomponents.Text("Create")),
	)
}
