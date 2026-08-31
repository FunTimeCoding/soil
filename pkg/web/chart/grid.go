package chart

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func (c *Chart) grid() []gomponents.Node {
	var nodes []gomponents.Node

	for _, fraction := range []float64{0, 0.25, 0.5, 0.75, 1} {
		value := c.maximum * fraction
		y := c.vertical(value)
		class := "chart-grid"

		if fraction == 1 {
			class = "chart-ceiling"
		}

		nodes = append(
			nodes,
			svgLine(
				constant.ChartMarginLeft,
				y,
				constant.ChartViewWidth-constant.ChartMarginRight,
				y,
				class,
			),
			svgText(
				constant.ChartMarginLeft-6,
				y+3,
				"chart-label chart-label-value",
				fmt.Sprintf("%.0f%%", value),
			),
		)
	}

	return nodes
}
