package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func layout(content ...gomponents.Node) gomponents.Node {
	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			gomponents.Attr("data-theme", "dark"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(
					html.Name("viewport"),
					html.Content("width=device-width, initial-scale=1"),
				),
				html.TitleEl(gomponents.Text("conversations")),
				html.Link(html.Rel("stylesheet"), html.Href(webConstant.Pico)),
				html.Script(html.Src(webConstant.Extended)),
				html.StyleEl(gomponents.Raw(webConstant.ThemeHearth)),
				html.StyleEl(gomponents.Raw(constant.ConversationStyle)),
			),
			html.Body(
				gomponents.Group(content),
				html.Script(gomponents.Raw(constant.InfiniteScrollScript)),
				html.Script(gomponents.Raw(constant.ScrollToBottomScript)),
				html.Script(gomponents.Raw(constant.HitNavigationScript)),
			),
		),
	)
}
