package board_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) strip(v *layout.Entry) gomponents.Node {
	rows := p.Rows(v)

	if rows == nil {
		return nil
	}

	return html.Div(
		html.Class("board-rows"),
		extended.StreamSwap(EventName(v.Label)),
		rows,
	)
}
