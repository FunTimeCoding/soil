package web

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (s *Server) recentVersions() gomponents.Node {
	since := time.Now().Add(-7 * 24 * time.Hour).UTC().Format(time.RFC3339)
	versions, e := s.service.RecentVersions(since, 10)

	if e != nil || len(versions) == 0 {
		return gomponents.Group(
			[]gomponents.Node{
				html.H3(gomponents.Text("Recent Changes")),
				html.P(gomponents.Text("No recent changes.")),
			},
		)
	}

	var rows []gomponents.Node

	for _, v := range versions {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(
					library.Parse(time.RFC3339, v.ChangedAt),
				),
				html.Td(gomponents.Text(v.ChangeType)),
				html.Td(
					html.A(
						gomponents.Attr(
							"href",
							fmt.Sprintf(
								"/memories/%d",
								v.MemoryIdentifier,
							),
						),
						gomponents.Text(v.Name),
					),
				),
				html.Td(html.Small(gomponents.Text(v.Source))),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Recent Changes")),
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Action")),
						html.Th(gomponents.Text("Memory")),
						html.Th(gomponents.Text("Source")),
					),
				),
				html.TBody(rows...),
			),
		},
	)
}
