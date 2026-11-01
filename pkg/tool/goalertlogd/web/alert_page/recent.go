package alert_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Recent(records []record.Record) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(records))

	for _, r := range records {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(html.Href(Path(r.Name)), gomponents.Text(r.Name)),
				),
				html.Td(component.SeverityBadge(r.Severity)),
				html.Td(gomponents.Text(component.Truncate(r.Summary, 80))),
				layout.TimeCell(r.Start),
				component.EndCell(r.End),
				html.Td(component.StatusBadge(r.End)),
			),
		)
	}

	return table.New("Name", "Severity", "Summary", "Start", "End", "Status").
		Rows(rows).
		Empty("No alerts in this time range.").
		Node()
}
