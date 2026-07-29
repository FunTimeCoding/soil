package layout

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func TimeCell(t time.Time) gomponents.Node {
	return html.Td(
		html.Class(constant.LayoutTimeCellClass),
		html.Small(gomponents.Text(library.FormatCompact(t))),
	)
}
