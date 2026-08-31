package chart

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (c *Chart) legend() gomponents.Node {
	var items []gomponents.Node

	for _, s := range c.series {
		items = append(
			items,
			html.Span(
				html.Class("chart-legend-item"),
				html.Span(
					html.Class(fmt.Sprintf("chart-legend-dot %s", s.Class)),
				),
				gomponents.Text(s.Label),
			),
		)
	}

	return html.Div(html.Class("chart-legend"), gomponents.Group(items))
}
