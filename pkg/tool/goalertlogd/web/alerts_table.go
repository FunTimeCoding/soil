package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func alertsTable(records []record.Record) gomponents.Node {
	if len(records) == 0 {
		return html.P(html.Em(gomponents.Text("No occurrences found.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Fingerprint")),
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
						html.Td(gomponents.Text(r.Fingerprint)),
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
