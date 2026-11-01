package user_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func PasswordForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.SetPasswordPath, "Set").
		Identifier(constant.PasswordForm).
		Heading("Set password").
		Alert(alert).
		Text(
			constant.AccountField,
			"Account",
			values.Get(constant.AccountField),
		).
		Required().
		Password(constant.PasswordField, "New password").
		Required().
		Node()
}
