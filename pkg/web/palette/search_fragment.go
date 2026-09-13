package palette

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func SearchFragment(
	endpoint string,
	placeholder string,
) gomponents.Node {
	return html.Div(
		html.Class("palette-body"),
		html.Input(
			html.ID("palette-search-input"),
			html.Type("text"),
			html.Class("palette-input"),
			gomponents.Attr("placeholder", placeholder),
			gomponents.Attr("autocomplete", "off"),
			extended.Get(endpoint),
			extended.Trigger(constant.TriggerType),
			extended.Target("#palette-search-results"),
			extended.Swap(constant.SwapOuter),
			gomponents.Attr("name", "q"),
			gomponents.Attr("autofocus", ""),
		),
		html.Div(html.ID("palette-search-results")),
	)
}
