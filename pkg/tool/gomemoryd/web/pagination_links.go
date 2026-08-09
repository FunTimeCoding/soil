package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func paginationLinks(
	page int,
	offset int,
	limit int,
	total int,
	tag string,
	memoryType string,
	scope string,
) []gomponents.Node {
	var navigation []gomponents.Node

	if page > 1 {
		navigation = append(
			navigation,
			html.A(
				gomponents.Attr(
					"href",
					pageLink(page-1, tag, memoryType, scope),
				),
				gomponents.Text("← Newer"),
			),
		)
	}

	if offset+limit < total {
		if len(navigation) > 0 {
			navigation = append(navigation, gomponents.Text(" · "))
		}

		navigation = append(
			navigation,
			html.A(
				gomponents.Attr(
					"href",
					pageLink(page+1, tag, memoryType, scope),
				),
				gomponents.Text("Older →"),
			),
		)
	}

	if len(navigation) == 0 {
		return nil
	}

	return []gomponents.Node{html.P(gomponents.Group(navigation))}
}
