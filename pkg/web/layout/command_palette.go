package layout

import (
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
				gomponents.Attr("hx-get", endpoint),
				gomponents.Attr("hx-trigger", "keyup changed delay:200ms"),
				gomponents.Attr("hx-target", "#palette-results"),
				gomponents.Attr("hx-swap", "outerHTML"),
				gomponents.Attr("name", "q"),
			),
			html.Div(
				html.ID("palette-results"),
				gomponents.Attr("hx-get", endpoint),
				gomponents.Attr("hx-trigger", "load"),
				gomponents.Attr("hx-swap", "outerHTML"),
			),
		),
	)
}
