package board_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Rows(v *layout.Entry) gomponents.Node {
	labels := v.RowLabels()

	if len(labels) == 0 {
		return nil
	}

	values := p.values[v.Label]
	var nodes []gomponents.Node

	for i, label := range labels {
		value := constant.PendingValue

		if i < len(values) {
			value = values[i]
		}

		nodes = append(
			nodes,
			html.Span(
				gomponents.Text(label),
				gomponents.Text(" "),
				html.Span(html.Class("board-row-value"), gomponents.Text(value)),
			),
		)
	}

	return gomponents.Group(nodes)
}
