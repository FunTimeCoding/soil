package impression_page

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func listTable(impressions []record.Impression) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(impressions))

	for _, i := range impressions {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(library.Parse(time.RFC3339, i.CreatedAt)),
				html.Td(html.Em(gomponents.Text(i.Source))),
				html.Td(gomponents.Text(i.Content)),
			),
		)
	}

	return table.New("Time", "Source", "Content").
		Rows(rows).
		Empty("No impressions.").
		Node()
}
