package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func paginationControls(
	offset, total int,
	startValue, endValue string,
) gomponents.Node {
	var nodes []gomponents.Node

	if offset > 0 {
		previous := offset - constant.PageSize

		if previous < 0 {
			previous = 0
		}

		nodes = append(
			nodes,
			html.A(
				html.Href(paginationLink(previous, startValue, endValue)),
				gomponents.Text("Previous"),
			),
		)
	}

	nodes = append(
		nodes,
		gomponents.Textf(
			" %d–%d of %d ",
			offset+1,
			min(offset+constant.PageSize, total),
			total,
		),
	)

	if offset+constant.PageSize < total {
		nodes = append(
			nodes,
			html.A(
				html.Href(
					paginationLink(
						offset+constant.PageSize,
						startValue,
						endValue,
					),
				),
				gomponents.Text("Next"),
			),
		)
	}

	return html.P(html.Class("pagination"), gomponents.Group(nodes))
}
