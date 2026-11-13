package layout

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) renderBody() []gomponents.Node {
	var body []gomponents.Node
	brand := p.identity.Name()

	if brand != "" {
		cell := []gomponents.Node{gomponents.Attr("href", constant.RootPath)}

		if p.liveEndpoint != "" {
			cell = append(cell, connectionDot())
		}

		node := html.A(append(cell, html.Strong(gomponents.Text(brand)))...)
		var navigation []gomponents.Node

		for _, item := range p.items {
			navigation = append(navigation, item.Render(p.path))
		}

		navigation = append(navigation, p.navigation...)

		if p.HasCommandPalette() {
			navigation = append(
				navigation,
				html.Li(
					html.A(
						gomponents.Attr("href", "#"),
						gomponents.Attr(
							"onclick",
							"document.getElementById('palette').showModal(); return false",
						),
						gomponents.Text("⌘K"),
					),
				),
			)
		}

		body = append(
			body,
			html.Nav(
				p.container(),
				html.Ul(html.Li(node)),
				html.Ul(navigation...),
			),
		)
	}

	body = append(
		body,
		html.Div(html.ID(constant.LayoutNotificationRegion), p.container()),
		html.Div(html.ID(constant.LayoutStatusLine), p.container()),
	)
	var mainAttrs []gomponents.Node
	mainAttrs = append(mainAttrs, p.container())

	if p.liveEndpoint != "" {
		endpoint := p.liveEndpoint

		if p.liveParams != "" {
			endpoint = fmt.Sprintf("%s?%s", endpoint, p.liveParams)
		}

		mainAttrs = append(
			mainAttrs,
			extended.Extension("sse"),
			extended.Connect(endpoint),
		)
	}

	if len(p.summary) > 0 {
		swap := summaryStrip(p.summary)

		if p.liveEndpoint != "" {
			swap = summaryStripLive(p.summary)
		}

		mainAttrs = append(mainAttrs, swap)
	}

	mainAttrs = append(mainAttrs, gomponents.Group(p.content))
	body = append(body, html.Main(mainAttrs...))

	if p.HasCommandPalette() {
		body = append(
			body,
			commandPalette(constant.PalettePath),
			html.Script(gomponents.Raw(constant.PaletteScript)),
		)
	}

	body = append(
		body,
		html.Script(gomponents.Raw(constant.NotificationScript)),
	)

	if p.liveEndpoint != "" {
		body = append(
			body,
			html.Script(gomponents.Raw(constant.ConnectionScript)),
		)
	}

	body = append(body, p.footer...)

	return body
}
