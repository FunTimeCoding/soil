package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func documentsNavigation(
	collection string,
	page int,
	hasMore bool,
	sourceType string,
) gomponents.Node {
	var links []gomponents.Node
	basePath := fmt.Sprintf("/documents/%s", collection)

	if page > 1 {
		links = append(
			links,
			html.A(
				extended.Get(documentsLink(basePath, page-1, sourceType)),
				extended.Target("#documents-content"),
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
				extended.Get(documentsLink(basePath, page+1, sourceType)),
				extended.Target("#documents-content"),
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
