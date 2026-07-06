package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (s *Server) topTable() gomponents.Node {
	now := time.Now()
	records := s.store.MustTop(25, now.Add(-7*24*time.Hour), now)

	if len(records) == 0 {
		return html.P(html.Em(gomponents.Text("No alerts in the last 7 days.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Count")),
				html.Th(gomponents.Text("Avg Duration")),
				html.Th(gomponents.Text("Firing")),
				html.Th(gomponents.Text("Severity")),
			),
		),
		html.TBody(
			gomponents.Map(
				records,
				func(r store.TopRecord) gomponents.Node {
					return html.Tr(
						html.Td(
							html.A(
								html.Class("alert-link"),
								html.Href(
									fmt.Sprintf(
										"/alerts?name=%s",
										r.Name,
									),
								),
								gomponents.Text(r.Name),
							),
						),
						html.Td(gomponents.Textf("%d", r.Count)),
						html.Td(gomponents.Text(formatDuration(r.AverageDuration))),
						html.Td(gomponents.Textf("%d", r.CurrentlyFiring)),
						html.Td(severityBadge(r.Severity)),
					)
				},
			),
		),
	)
}
