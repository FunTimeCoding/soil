package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func memberForm() gomponents.Node {
	return html.Form(
		html.Method("post"),
		html.Action(constant.AddMemberPath),
		html.H2(gomponents.Text("Membership")),
		field(constant.GroupField, "Group", constant.TextInputType, true),
		field(constant.AccountField, "Account", constant.TextInputType, true),
		html.Button(html.Type("submit"), gomponents.Text("Add")),
		html.Button(
			html.Type("submit"),
			html.FormAction(constant.RemoveMemberPath),
			gomponents.Text("Remove"),
		),
	)
}
