package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func recentTable(records []record.Record) gomponents.Node {
	if len(records) == 0 {
		return html.P(html.Em(gomponents.Text("No alerts in this time range.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Severity")),
				html.Th(gomponents.Text("Summary")),
				html.Th(gomponents.Text("Start")),
				html.Th(gomponents.Text("End")),
				html.Th(gomponents.Text("Status")),
			),
		),
		html.TBody(
			gomponents.Map(
				records,
				func(r record.Record) gomponents.Node {
					status := "firing"

					if r.End != nil {
						status = "resolved"
					}

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
						html.Td(severityBadge(r.Severity)),
						html.Td(gomponents.Text(truncate(r.Summary, 80))),
						html.Td(gomponents.Text(formatTime(r.Start))),
						html.Td(gomponents.Text(formatTimePointer(r.End))),
						html.Td(statusBadge(status)),
					)
				},
			),
		),
	)
}
