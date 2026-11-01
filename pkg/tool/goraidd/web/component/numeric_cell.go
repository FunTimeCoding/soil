package component

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func NumericCell(text string) gomponents.Node {
	return html.Td(html.Class(webConstant.NumericClass), gomponents.Text(text))
}
