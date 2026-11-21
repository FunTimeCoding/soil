package remediation_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/tile"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func count(
	n int,
	label string,
) gomponents.Node {
	return tile.New(
		html.Strong(gomponents.Text(strconv.Itoa(n))),
		html.Br(),
		html.Small(gomponents.Text(label)),
	).Node()
}
