package chart

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func (c *Chart) dayTicks() []gomponents.Node {
	var nodes []gomponents.Node
	day := time.Date(
		c.start.Year(),
		c.start.Month(),
		c.start.Day(),
		0,
		0,
		0,
		0,
		c.start.Location(),
	).AddDate(0, 0, 1)

	for ; day.Before(c.end); day = day.AddDate(0, 0, 1) {
		x := c.horizontal(day)
		nodes = append(
			nodes,
			svgLine(
				x,
				constant.ChartMarginTop,
				x,
				constant.ChartViewHeight-constant.ChartMarginBottom,
				"chart-grid",
			),
			svgText(
				x,
				constant.ChartViewHeight-constant.ChartMarginBottom+16,
				"chart-label chart-label-day",
				day.Format("Mon"),
			),
		)
	}

	return nodes
}
