package collection_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/status"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func List(collections []status.CollectionStatus) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(collections))

	for _, c := range collections {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(Path(c.Name, "", 1)),
						gomponents.Text(c.Name),
					),
				),
				html.Td(gomponents.Text(c.Type)),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", c.DocumentCount),
				),
				html.Td(html.Small(gomponents.Text(c.Path))),
				html.Td(html.Small(gomponents.Text(c.Pattern))),
			),
		)
	}

	return table.New("Name", "Type", "Documents", "Path", "Pattern").
		Numeric(2).
		Rows(rows).
		Empty("No collections.").
		Node()
}
