package dial

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func needle(m *marker) gomponents.Node {
	return html.Div(
		html.Class(constant.DialNeedleClass),
		html.Style(
			join.Empty(
				constant.DialAngleVariable,
				": ",
				fmt.Sprintf("%.1fdeg", degrees(m.share)-constant.DialQuarter),
			),
		),
		html.Div(
			html.Class(m.class),
			html.Title(m.title),
			html.Style(fmt.Sprintf("--dial-color: %s;", m.color)),
		),
	)
}
