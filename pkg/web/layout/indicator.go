package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Indicator(
	identifier string,
	text string,
) gomponents.Node {
	return html.Div(
		html.ID(identifier),
		html.Class(constant.ExtendedIndicatorClass),
		html.Span(html.Class(constant.IndicatorMarkClass)),
		html.Small(gomponents.Text(text)),
	)
}
