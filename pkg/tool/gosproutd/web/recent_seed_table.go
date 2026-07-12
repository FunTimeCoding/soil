package web

import (
	"github.com/funtimecoding/soil/pkg/time"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) recentSeedTable() gomponents.Node {
	seeds := s.service.RecentSeeds()
	var rows []gomponents.Node

	for _, v := range seeds {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Class("time-cell"),
					html.Small(
						gomponents.Text(time.FormatCompact(v.ModifiedAt)),
					),
				),
				html.Td(gomponents.Text(v.Name)),
				html.Td(
					html.Style("color: var(--pico-muted-color);"),
					gomponents.Text(seedDirectory(v.Path)),
				),
			),
		)
	}

	return html.Table(
		html.Class("seed-table"),
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Modified")),
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Path")),
			),
		),
		html.TBody(gomponents.Group(rows)),
	)
}
