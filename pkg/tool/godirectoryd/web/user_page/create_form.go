package user_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func CreateForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.CreateUserPath, "Create").
		Identifier(constant.CreateUserForm).
		Heading("Create user").
		Alert(alert).
		Text(
			constant.AccountField,
			"Account",
			values.Get(constant.AccountField),
		).
		Required().
		Text(constant.NameField, "Name", values.Get(constant.NameField)).
		Required().
		Text(
			constant.SurnameField,
			"Surname",
			values.Get(constant.SurnameField),
		).
		Required().
		Mail(constant.MailField, "Mail", values.Get(constant.MailField)).
		Password(constant.PasswordField, "Password").
		Node()
}
