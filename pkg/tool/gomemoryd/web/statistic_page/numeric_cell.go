package statistic_page

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func numericCell(value int) gomponents.Node {
	return html.Td(
		html.Class(webConstant.NumericClass),
		gomponents.Textf("%d", value),
	)
}
