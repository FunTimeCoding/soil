package chart

import (
	"fmt"
	"maragu.dev/gomponents"
)

func svgLine(
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	class string,
) gomponents.Node {
	return gomponents.El(
		"line",
		gomponents.Attr("x1", fmt.Sprintf("%.1f", x1)),
		gomponents.Attr("y1", fmt.Sprintf("%.1f", y1)),
		gomponents.Attr("x2", fmt.Sprintf("%.1f", x2)),
		gomponents.Attr("y2", fmt.Sprintf("%.1f", y2)),
		gomponents.Attr("class", class),
	)
}
