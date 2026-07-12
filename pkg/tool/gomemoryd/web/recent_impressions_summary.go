package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (s *Server) recentImpressionsSummary() gomponents.Node {
	since := time.Now().Add(-48 * time.Hour).UTC().Format(time.RFC3339)
	impressions, e := s.service.RecentImpressions(since)

	if e != nil || len(impressions) == 0 {
		return gomponents.Group(
			[]gomponents.Node{
				html.H3(gomponents.Text("Recent Impressions")),
				html.P(gomponents.Text("No recent impressions.")),
			},
		)
	}

	var rows []gomponents.Node

	for _, i := range impressions {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Class(timeCellClass),
					html.Small(gomponents.Text(formatTime(i.CreatedAt))),
				),
				html.Td(html.Em(gomponents.Text(i.Source))),
				html.Td(gomponents.Text(i.Content)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text("Recent Impressions")),
			html.Table(
				html.THead(
					html.Tr(
						html.Th(gomponents.Text("Time")),
						html.Th(gomponents.Text("Source")),
						html.Th(gomponents.Text("Content")),
					),
				),
				html.TBody(rows...),
			),
		},
	)
}
