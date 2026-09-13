package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func userTable(v []*user.User) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(v))

	for _, u := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(u.Account)),
				html.Td(gomponents.Text(u.Name)),
				html.Td(gomponents.Text(u.Mail)),
				html.Td(
					html.Form(
						html.Method("post"),
						html.Action(constant.DeleteUserPath),
						html.Input(
							html.Type("hidden"),
							html.Name(constant.AccountField),
							html.Value(u.Account),
						),
						html.Button(
							html.Type("submit"),
							gomponents.Text("Delete"),
						),
					),
				),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Account")),
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Mail")),
				html.Th(gomponents.Text("")),
			),
		),
		html.TBody(rows...),
	)
}
