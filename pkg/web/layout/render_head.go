package layout

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) renderHead(title string) []gomponents.Node {
	var head []gomponents.Node
	head = append(
		head,
		html.Meta(html.Charset("utf-8")),
		html.Meta(
			html.Name("viewport"),
			html.Content("width=device-width, initial-scale=1"),
		),
		html.TitleEl(gomponents.Text(title)),
		html.Link(html.Rel("stylesheet"), html.Href(constant.Pico)),
		html.StyleEl(gomponents.Raw(constant.BaseStyle)),
		html.StyleEl(gomponents.Raw(constant.NotificationStyle)),
		html.StyleEl(gomponents.Raw(constant.IndicatorStyle)),
		html.Script(html.Src(p.extendedSource())),
	)

	for _, s := range p.script {
		head = append(head, html.Script(html.Src(s)))
	}

	if p.liveEndpoint != "" {
		head = append(
			head,
			html.Script(html.Src(p.serverSideSource())),
			html.StyleEl(gomponents.Raw(constant.ConnectionStyle)),
		)
	}

	if p.theme != "" {
		head = append(head, html.StyleEl(gomponents.Raw(p.theme)))
	}

	if p.paletteEndpoint != "" {
		head = append(head, html.StyleEl(gomponents.Raw(constant.PaletteStyle)))
	}

	if p.style != "" {
		head = append(head, html.StyleEl(gomponents.Raw(p.style)))
	}

	return head
}
