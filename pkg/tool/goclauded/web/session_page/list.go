package session_page

import (
	"github.com/funtimecoding/soil/pkg/strings/shorten"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(
	sessions []*enriched_session.Session,
	active bool,
	page int,
	total int,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(sessions))

	for _, s := range sessions {
		name := s.Alias

		if name == "" {
			name = s.Slug
		}

		if name == "" {
			name = shorten.Prefix(s.Identifier, 8)
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(html.Href(Path(s.Identifier)), gomponents.Text(name)),
					gomponents.If(
						s.Active,
						badge.New(constant.ActiveLabel).Node(),
					),
				),
				html.Td(html.Small(gomponents.Text(s.Description))),
				numeric(s.Lines),
				numeric(s.TurnCount),
			),
		)
	}

	all := chip.New("All").Link(ListPath(false, 1))
	only := chip.New("Active").Link(ListPath(true, 1))

	if active {
		only.Current()
	} else {
		all.Current()
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.Div(
				html.Class(webConstant.ClusterClass),
				all.Node(),
				only.Node(),
			),
			table.New("Name", "Description", "Lines", "Turns").
				Numeric(2, 3).
				Rows(rows).
				Empty("No sessions.").
				Node(),
			pages(active, page, total),
		},
	)
}
