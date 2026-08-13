package web

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func memoryLink(
	identifier int64,
	name string,
) gomponents.Node {
	return html.A(
		gomponents.Attr("href", fmt.Sprintf("/memories/%d", identifier)),
		gomponents.Text(name),
	)
}
