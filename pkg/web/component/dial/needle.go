package dial

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/component/dial/marker"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func needle(m *marker.Marker) gomponents.Node {
	return html.Div(
		html.Class(constant.DialNeedleClass),
		html.Style(
			join.Empty(
				constant.DialAngleVariable,
				": ",
				fmt.Sprintf("%.1fdeg", degrees(m.Share)-constant.DialQuarter),
			),
		),
		html.Div(
			html.Class(m.Class),
			html.Title(m.Title),
			html.Style(fmt.Sprintf("--dial-color: %s;", m.Color)),
		),
	)
}
