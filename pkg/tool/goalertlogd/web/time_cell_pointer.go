package web

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func timeCellPointer(t *time.Time) gomponents.Node {
	if t == nil {
		return html.Td(html.Class(layout.TimeCellClass))
	}

	return layout.TimeCell(*t)
}
