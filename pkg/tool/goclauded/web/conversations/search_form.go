package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func searchForm() gomponents.Node {
	return html.Form(
		html.Class("search-form"),
		extended.Get("/conversations/search"),
		extended.Target("#sidebar-entries"),
		extended.Trigger(web.TriggerForm),
		extended.Sync(web.SyncReplace),
		html.Input(
			html.Type(web.SearchInputType),
			html.Name(constant.QueryField),
			html.Class("search-input"),
			gomponents.Attr("placeholder", "Search conversations..."),
			gomponents.Attr("autocomplete", "off"),
		),
		html.Div(
			html.Class("search-kinds"),
			kindToggle(constant.BlockMessage, true),
			kindToggle(constant.BlockEdit, false),
			kindToggle(constant.BlockCall, false),
		),
	)
}
