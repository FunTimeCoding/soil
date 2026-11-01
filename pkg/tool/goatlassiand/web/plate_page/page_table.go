package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func pageTable(pages []*page.Page) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(pages))

	for _, p := range pages {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.A(
						html.Href(p.Link),
						html.Target("_blank"),
						gomponents.Text(p.Name),
					),
				),
				layout.TimeCell(p.Raw.Version.CreatedAt),
			),
		)
	}

	return table.New().
		Headerless().
		Class(constant.PageTableClass).
		Rows(rows).
		Node()
}
