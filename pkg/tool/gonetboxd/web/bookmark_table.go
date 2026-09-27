package web

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/object_type"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func bookmarkTable(bookmarks []*bookmark.Bookmark) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(bookmarks))

	for _, b := range bookmarks {
		rows = append(
			rows,
			html.Tr(
				html.Td(bookmarkLink(b)),
				html.Td(
					html.Class("reference"),
					gomponents.Text(object_type.Label(b.ObjectType)),
				),
				layout.TimeCell(b.Created),
			),
		)
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Type")),
				html.Th(gomponents.Text("Bookmarked")),
			),
		),
		html.TBody(rows...),
	)
}
