package board_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/request_entry"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func Requests(
	requests []*request_entry.Entry,
	multiple bool,
) gomponents.Node {
	if len(requests) == 0 {
		return gomponents.Group(nil)
	}

	var columns []string

	if multiple {
		columns = append(columns, "Instance")
	}

	columns = append(columns, "Request", "Title", "State", "Created")
	rows := make([]gomponents.Node, 0, len(requests))

	for _, entry := range requests {
		r := entry.Request
		var created time.Time

		if r.Create != nil {
			created = *r.Create
		}

		rows = append(
			rows,
			html.Tr(
				gomponents.If(
					multiple,
					html.Td(
						html.Class(constant.ReferenceClass),
						gomponents.Text(entry.Instance),
					),
				),
				html.Td(
					html.A(
						html.Href(r.Link),
						html.Target("_blank"),
						gomponents.Text(fmt.Sprintf("!%d", r.Identifier)),
					),
				),
				html.Td(gomponents.Text(r.Title)),
				html.Td(
					html.Class(constant.ReferenceClass),
					gomponents.Text(r.State),
				),
				layout.TimeCell(created),
			),
		)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(constant.RequestTitle)),
			table.New(columns...).Class(constant.BoardClass).Rows(rows).Node(),
		},
	)
}
