package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) seedTable() gomponents.Node {
	seeds := s.service.Seeds()
	var rows []gomponents.Node

	for _, v := range seeds {
		rows = append(
			rows,
			html.Tr(
				html.Td(html.Class("drag-handle"), gomponents.Text("⠿")),
				html.Td(
					html.Class("position-cell"),
					gomponents.Text(fmt.Sprintf("%d", v.Position)),
				),
				html.Td(gomponents.Text(v.Name)),
				html.Td(
					html.Style("color: var(--pico-muted-color);"),
					gomponents.Text(seedDirectory(v.Path)),
				),
				html.Input(
					html.Type("hidden"),
					html.Name("item"),
					html.Value(fmt.Sprintf("%d", v.Identifier)),
				),
			),
		)
	}

	return html.Table(
		html.Class("seed-table"),
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("")),
				html.Th(html.Class("position-cell"), gomponents.Text("#")),
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Path")),
			),
		),
		html.TBody(html.Class("sortable"), gomponents.Group(rows)),
	)
}
