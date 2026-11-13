package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func OptionalTimeCell(
	t *time.Time,
	absent string,
) gomponents.Node {
	if t != nil {
		return TimeCell(*t)
	}

	return html.Td(
		html.Class(constant.LayoutTimeCellClass),
		gomponents.If(
			absent != "",
			html.Small(
				html.Class(constant.AbsentClass),
				gomponents.Text(absent),
			),
		),
	)
}
