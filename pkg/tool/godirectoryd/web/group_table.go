package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func groupTable(v []*group.Group) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(v))

	for _, g := range v {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(g.Name)),
				html.Td(gomponents.Text(strconv.Itoa(g.Number))),
				html.Td(gomponents.Text(join.CommaSpace(g.Member))),
				html.Td(
					html.Form(
						html.Method("post"),
						html.Action(constant.DeleteGroupPath),
						html.Input(
							html.Type("hidden"),
							html.Name(constant.NameField),
							html.Value(g.Name),
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
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Number")),
				html.Th(gomponents.Text("Members")),
				html.Th(gomponents.Text("")),
			),
		),
		html.TBody(rows...),
	)
}
