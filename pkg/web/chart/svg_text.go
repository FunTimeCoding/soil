package chart

import (
	"fmt"
	"maragu.dev/gomponents"
)

func svgText(
	x float64,
	y float64,
	class string,
	value string,
) gomponents.Node {
	return gomponents.El(
		"text",
		gomponents.Attr("x", fmt.Sprintf("%.1f", x)),
		gomponents.Attr("y", fmt.Sprintf("%.1f", y)),
		gomponents.Attr("class", class),
		gomponents.Text(value),
	)
}
