package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sortLinks(modified bool) gomponents.Node {
	priority := gomponents.Node(html.Strong(gomponents.Text("priority")))
	recent := gomponents.Node(
		html.A(
			gomponents.Attr(
				"href",
				fmt.Sprintf(
					"%s?%s=%s",
					constant.DashboardPath,
					constant.SortParameter,
					constant.SortModified,
				),
			),
			gomponents.Text(constant.SortModified),
		),
	)

	if modified {
		priority = html.A(
			gomponents.Attr("href", constant.DashboardPath),
			gomponents.Text("priority"),
		)
		recent = html.Strong(gomponents.Text(constant.SortModified))
	}

	return html.P(
		html.Small(priority, gomponents.Text(" · "), recent),
	)
}
