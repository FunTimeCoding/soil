package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func commandPalette(endpoint string) gomponents.Node {
	return html.Dialog(
		html.ID("palette"),
		html.Class("palette-dialog"),
		html.Div(
			html.Class("palette-body"),
			html.Input(
				html.ID("palette-input"),
				html.Type("text"),
				html.Class("palette-input"),
				gomponents.Attr("placeholder", "Type a command..."),
				gomponents.Attr("autocomplete", "off"),
				extended.Get(endpoint),
				extended.Trigger(constant.TriggerType),
				extended.Target("#palette-results"),
				extended.Swap(constant.SwapOuter),
				gomponents.Attr("name", "q"),
			),
			html.Div(
				html.ID("palette-results"),
				extended.Get(endpoint),
				extended.Trigger(constant.TriggerLoad),
				extended.Swap(constant.SwapOuter),
			),
		),
	)
}
