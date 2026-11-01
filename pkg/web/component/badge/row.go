package badge

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Row(badges ...gomponents.Node) gomponents.Node {
	if len(badges) == 0 {
		return gomponents.Group(nil)
	}

	return html.Div(
		html.Class(constant.BadgeRowClass),
		gomponents.Group(badges),
	)
}
