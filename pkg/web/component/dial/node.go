package dial

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (d *Dial) Node() gomponents.Node {
	needles := make([]gomponents.Node, 0, len(d.markers))

	for _, m := range d.markers {
		needles = append(needles, needle(m))
	}

	return html.Div(
		html.Class(constant.DialClass),
		html.Style(
			join.Empty(constant.DialBackgroundVariable, ": ", d.background()),
		),
		gomponents.Group(needles),
		html.Div(
			html.Class(constant.DialCenterClass),
			gomponents.Group(d.center),
		),
	)
}
