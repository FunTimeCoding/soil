package meter

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Row(meters ...gomponents.Node) gomponents.Node {
	if len(meters) == 0 {
		return gomponents.Group(nil)
	}

	return html.Div(
		html.Class(constant.MeterRowClass),
		gomponents.Group(meters),
	)
}
