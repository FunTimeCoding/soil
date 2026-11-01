package certificate_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(records []record.Record) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(records))

	for _, r := range records {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(r.CommonName)),
				html.Td(gomponents.Text(r.Kind)),
				html.Td(gomponents.Text(r.Issuer)),
				layout.TimeCell(r.End),
				html.Td(stateBadge(&r)),
			),
		)
	}

	return table.New("Common name", "Kind", "Issuer", "Expires", "State").
		Rows(rows).
		Empty("Nothing issued yet.").
		Node()
}
