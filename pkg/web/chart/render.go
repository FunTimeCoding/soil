package chart

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (c *Chart) Render() gomponents.Node {
	nodes := append(c.grid(), c.dayTicks()...)

	if c.guide {
		nodes = append(
			nodes,
			svgLine(
				c.horizontal(c.start),
				c.vertical(0),
				c.horizontal(c.end),
				c.vertical(c.maximum),
				"chart-guide",
			),
		)
	}

	if !c.now.IsZero() {
		nodes = append(
			nodes,
			svgLine(
				c.horizontal(c.now),
				constant.ChartMarginTop,
				c.horizontal(c.now),
				constant.ChartViewHeight-constant.ChartMarginBottom,
				"chart-now",
			),
		)
	}

	for _, s := range c.series {
		nodes = append(
			nodes,
			gomponents.El(
				"path",
				gomponents.Attr("d", c.stepPath(s)),
				gomponents.Attr("class", fmt.Sprintf("chart-line %s", s.Class)),
			),
		)

		if c.projection {
			nodes = append(nodes, c.projectionLine(s)...)
		}
	}

	return html.Div(
		html.Class("chart"),
		html.SVG(
			gomponents.Attr(
				"viewBox",
				fmt.Sprintf(
					"0 0 %.0f %.0f",
					constant.ChartViewWidth,
					constant.ChartViewHeight,
				),
			),
			gomponents.Group(nodes),
		),
		c.legend(),
	)
}
