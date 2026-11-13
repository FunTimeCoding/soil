package alert_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Occurrences(
	name string,
	records []record.Record,
) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(records))

	for _, r := range records {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(r.Fingerprint)),
				html.Td(component.SeverityBadge(r.Severity)),
				layout.ClampCell(r.Summary),
				layout.TimeCell(r.Start),
				layout.OptionalTimeCell(r.End, ""),
				html.Td(component.StatusBadge(r.End)),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H1(gomponents.Textf("Alert: %s", name)),
			html.P(gomponents.Textf("%d occurrences", len(records))),
			table.New(
				"Fingerprint",
				"Severity",
				"Summary",
				"Start",
				"End",
				"Status",
			).
				Rows(rows).
				Empty("No occurrences found.").
				Node(),
		},
	)
}
