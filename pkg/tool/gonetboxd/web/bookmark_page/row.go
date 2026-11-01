package bookmark_page

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"github.com/funtimecoding/soil/pkg/netbox/object_type"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func row(b *bookmark.Bookmark) gomponents.Node {
	return html.Tr(
		html.Td(link(b)),
		html.Td(gomponents.Text(object_type.Label(b.ObjectType))),
		layout.TimeCell(b.Created),
	)
}
