package dashboard_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/top"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/alert_page"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Top(records []top.Record) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(records))

	for _, r := range records {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(alert_page.Path(r.Name)),
						gomponents.Text(r.Name),
					),
				),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", r.Count),
				),
				html.Td(gomponents.Text(duration(r.AverageDuration))),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", r.CurrentlyFiring),
				),
				html.Td(component.SeverityBadge(r.Severity)),
			),
		)
	}

	return table.New("Name", "Count", "Avg Duration", "Firing", "Severity").
		Numeric(1, 3).
		Rows(rows).
		Empty("No alerts in the last 7 days.").
		Node()
}
