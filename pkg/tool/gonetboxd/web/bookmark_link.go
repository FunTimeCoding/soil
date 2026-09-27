package web

import (
	"github.com/funtimecoding/soil/pkg/netbox/bookmark"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func bookmarkLink(b *bookmark.Bookmark) gomponents.Node {
	if b.Link == "" {
		return gomponents.Text(b.Display)
	}

	return html.A(
		html.Href(b.Link),
		html.Target("_blank"),
		gomponents.Text(b.Display),
	)
}
