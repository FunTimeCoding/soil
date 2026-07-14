package layout

import (
	library "github.com/funtimecoding/soil/pkg/time"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func TimeCell(t time.Time) gomponents.Node {
	return html.Td(
		html.Class(TimeCellClass),
		html.Small(gomponents.Text(library.FormatCompact(t))),
	)
}
