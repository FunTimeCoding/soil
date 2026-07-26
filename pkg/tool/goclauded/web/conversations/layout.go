package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	theme "github.com/funtimecoding/soil/pkg/web/theme/constant"
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
				html.Link(html.Rel("stylesheet"), html.Href(web.Pico)),
				html.Script(html.Src(web.Extended)),
				html.StyleEl(gomponents.Raw(theme.Hearth)),
				html.StyleEl(gomponents.Raw(constant.ConversationStyle)),
			),
			html.Body(
				gomponents.Group(content),
				html.Script(gomponents.Raw(constant.SidebarFilterScript)),
				html.Script(gomponents.Raw(constant.InfiniteScrollScript)),
				html.Script(gomponents.Raw(constant.ScrollToBottomScript)),
			),
		),
	)
}
