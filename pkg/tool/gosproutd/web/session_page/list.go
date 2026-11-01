package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/session"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(sessions []*session.Session) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(sessions))

	for _, v := range sessions {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(html.Href(Path(v.Name)), gomponents.Text(v.Name)),
				),
				count(v.Open),
				count(v.Pending),
				count(v.Cleared),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.If(
						v.Bumped > 0,
						html.Strong(gomponents.Textf("%d", v.Bumped)),
					),
				),
				layout.TimeCell(v.LastAt),
			),
		)
	}

	return table.New(
		"Session",
		"Open",
		"Waiting on me",
		"Cleared",
		"Bumped",
		"Last",
	).
		Numeric(1, 2, 3, 4).
		Rows(rows).
		Empty("No session has pushed a decision.").
		Node()
}
