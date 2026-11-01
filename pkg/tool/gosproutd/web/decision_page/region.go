package decision_page

import (
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func region(
	identifier uint,
	part string,
	content gomponents.Node,
) gomponents.Node {
	return html.Div(extended.StreamSwap(Event(identifier, part)), content)
}
