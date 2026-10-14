package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func searchForm() gomponents.Node {
	return html.Form(
		html.Class("search-form"),
		extended.Get("/conversations/search"),
		extended.Target("#sidebar-entries"),
		extended.Trigger(webConstant.TriggerForm),
		extended.Sync(webConstant.SyncReplace),
		html.Input(
			html.Type(webConstant.SearchInputType),
			html.Name(constant.QueryField),
			html.Class("search-input"),
			gomponents.Attr("placeholder", "Search conversations..."),
			gomponents.Attr("autocomplete", constant.AutocompleteOff),
		),
		html.Div(
			html.Class("search-kinds"),
			kindToggle(constant.BlockMessage, true),
			kindToggle(constant.BlockEdit, false),
			kindToggle(constant.BlockCall, false),
		),
	)
}
