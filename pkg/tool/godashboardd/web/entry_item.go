package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func entryItem(
	v *board.Entry,
	values map[string][]string,
) gomponents.Node {
	nodes := []gomponents.Node{
		html.A(
			html.Href(v.Link),
			html.Target("_blank"),
			html.Class("board-entry-link"),
			gomponents.Attr(constant.LabelAttribute, v.Label),
			html.Img(
				html.Src(iconLink(v.Icon)),
				html.Class("board-icon"),
				html.Alt(v.Label),
				html.Loading("lazy"),
			),
			html.Span(gomponents.Text(v.Label)),
		),
	}

	if strip := rowStrip(v, values[v.Label]); strip != nil {
		nodes = append(nodes, strip)
	}

	return html.Div(gomponents.Group(nodes))
}
