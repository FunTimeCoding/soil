package board_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func grid(
	columns int,
	class string,
	nodes []gomponents.Node,
) gomponents.Node {
	return html.Div(
		html.Class(class),
		html.Style(
			fmt.Sprintf("%s: %d", constant.BoardColumnsVariable, columns),
		),
		gomponents.Group(nodes),
	)
}
