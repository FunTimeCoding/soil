package web

import (
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func historyNavigation(
	page int,
	hasMore bool,
	kinds []string,
) gomponents.Node {
	var links []gomponents.Node

	if page > 1 {
		links = append(
			links,
			html.A(
				extended.Get(historyLink(page-1, kinds)),
				extended.Target("#history-content"),
				extended.Swap("innerHTML"),
				gomponents.Text("← Newer"),
			),
		)
	}

	if hasMore {
		if len(links) > 0 {
			links = append(links, gomponents.Text(" · "))
		}

		links = append(
			links,
			html.A(
				extended.Get(historyLink(page+1, kinds)),
				extended.Target("#history-content"),
				extended.Swap("innerHTML"),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(links) == 0 {
		return gomponents.Group(nil)
	}

	return html.P(gomponents.Group(links))
}
