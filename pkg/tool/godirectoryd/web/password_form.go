package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func passwordForm() gomponents.Node {
	return html.Form(
		html.Method("post"),
		html.Action(constant.SetPasswordPath),
		html.H2(gomponents.Text("Set password")),
		field(constant.AccountField, "Account", constant.TextInputType, true),
		field(
			constant.PasswordField,
			"New password",
			constant.PasswordInputType,
			true,
		),
		html.Button(html.Type("submit"), gomponents.Text("Set")),
	)
}
