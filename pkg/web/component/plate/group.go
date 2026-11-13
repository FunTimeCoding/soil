package plate

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Group(plates ...gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class(constant.PlateGroupClass),
		gomponents.Group(plates),
	)
}
