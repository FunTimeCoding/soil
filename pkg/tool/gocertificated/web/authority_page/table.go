package authority_page

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
				html.Td(gomponents.Text(r.Name)),
				html.Td(gomponents.Text(r.Kind)),
				html.Td(gomponents.Text(r.CommonName)),
				html.Td(gomponents.Text(permitted(&r))),
				layout.TimeCell(r.End),
				html.Td(publishedBadge(&r)),
			),
		)
	}

	return table.New(
		"Name",
		"Kind",
		"Common name",
		"Permits",
		"Expires",
		"Published",
	).
		Rows(rows).
		Empty("No authority yet. Create the root to begin the chain.").
		Node()
}
