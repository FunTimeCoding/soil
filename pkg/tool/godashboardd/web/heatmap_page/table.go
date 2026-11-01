package heatmap_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/types/summary"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(summaries []summary.Summary) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(summaries))

	for _, v := range summaries {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(v.Label)),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Text(fmt.Sprintf("%d", v.Count)),
				),
				layout.TimeCell(v.Last),
			),
		)
	}

	return table.New("Link", "Clicks", "Last clicked").
		Numeric(1).
		Rows(rows).
		Empty("No clicks recorded yet.").
		Node()
}
