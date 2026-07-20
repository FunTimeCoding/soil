package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/enriched_session"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sessionRow(s *enriched_session.Session) gomponents.Node {
	name := s.Slug

	if s.Alias != "" {
		name = s.Alias
	}

	if name == "" {
		name = s.Identifier[:8]
	}

	var label []gomponents.Node
	label = append(label, gomponents.Text(name))

	if s.Active {
		label = append(
			label,
			gomponents.Text(" "),
			html.Small(gomponents.Text("(active)")),
		)
	}

	return html.Tr(
		html.Td(
			html.A(
				gomponents.Attr(
					"href",
					fmt.Sprintf("/sessions/%s", s.Identifier),
				),
				gomponents.Group(label),
			),
		),
		html.Td(html.Small(gomponents.Text(s.Description))),
		html.Td(gomponents.Textf("%d", s.Lines)),
		html.Td(gomponents.Textf("%d", s.TurnCount)),
	)
}
