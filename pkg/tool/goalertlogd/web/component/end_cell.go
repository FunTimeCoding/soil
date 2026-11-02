package component

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func EndCell(t *time.Time) gomponents.Node {
	if t == nil {
		return html.Td(html.Class(webConstant.LayoutTimeCellClass))
	}

	return layout.TimeCell(*t)
}
