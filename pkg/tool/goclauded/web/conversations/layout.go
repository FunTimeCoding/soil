package conversations

import (
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/theme/constant"
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
				html.StyleEl(gomponents.Raw(constant.Hearth)),
				html.StyleEl(gomponents.Raw(inlineCSS)),
			),
			html.Body(
				gomponents.Group(content),
				html.Script(gomponents.Raw(filterJS)),
				html.Script(gomponents.Raw(infiniteScrollJS)),
				html.Script(gomponents.Raw(scrollToBottomJS)),
			),
		),
	)
}
