package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func boardGrid(
	columns int,
	class string,
	nodes []gomponents.Node,
) gomponents.Node {
	return html.Div(
		html.Class(class),
		html.Style(
			fmt.Sprintf("grid-template-columns: repeat(%d, 1fr);", columns),
		),
		gomponents.Group(nodes),
	)
}
