package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/permission"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Region(r *permission.Report) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(r.Rows))
	var details []gomponents.Node

	for _, w := range r.Rows {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(w.Blocker)),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", w.Prompts),
				),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", w.Sessions),
				),
				html.Td(gomponents.Text(reason(w.Reason))),
				layout.TimeCell(w.Last),
			),
		)

		if len(w.Commands) == 0 {
			continue
		}

		commands := make([]gomponents.Node, 0, len(w.Commands))

		for _, c := range w.Commands {
			commands = append(
				commands,
				html.Pre(gomponents.Text(commandText(c))),
			)
		}

		details = append(
			details,
			html.Details(
				html.Summary(gomponents.Text(w.Blocker)),
				gomponents.Group(commands),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			table.New(
				"Blocker",
				"Prompts",
				constant.SessionsTitle,
				"Reason",
				"Last",
			).
				Numeric(1, 2).
				Rows(rows).
				Empty("No permission prompts in this window.").
				Node(),
			gomponents.Group(details),
		},
	)
}
