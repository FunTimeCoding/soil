package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) searchResults(query string) gomponents.Node {
	if query == "" {
		return html.Div()
	}

	results, e := s.service.SearchMemories(
		PrefixQuery(query),
		constant.SearchLimit,
		"",
		"",
		constant.AllScope,
	)

	if e != nil {
		return html.Div(
			html.P(gomponents.Textf("Search failed: %s", e.Error())),
		)
	}

	if len(results) == 0 {
		return html.Div(html.P(gomponents.Text("No results.")))
	}

	var rows []gomponents.Node

	for _, m := range results {
		var pips []gomponents.Node

		for _, t := range m.Tags {
			pips = append(
				pips,
				html.Span(html.Class("tag-pip"), gomponents.Text(t)),
				gomponents.Text(" "),
			)
		}

		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf("/memories/%d", m.Identifier),
						),
						gomponents.Text(m.Name),
					),
				),
				html.Td(html.Small(gomponents.Text(m.Description))),
				html.Td(gomponents.Text(m.Type)),
				html.Td(gomponents.Group(pips)),
			),
		)
	}

	return html.Div(
		html.P(gomponents.Textf("%d results", len(results))),
		html.Table(
			html.THead(
				html.Tr(
					html.Th(gomponents.Text("Name")),
					html.Th(gomponents.Text("Description")),
					html.Th(gomponents.Text("Type")),
					html.Th(gomponents.Text("Tags")),
				),
			),
			html.TBody(rows...),
		),
	)
}
