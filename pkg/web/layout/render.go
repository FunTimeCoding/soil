package layout

import (
	"fmt"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Render() gomponents.Node {
	title := p.identity.Name()

	if p.title != "" {
		title = fmt.Sprintf("%s - %s", title, p.title)
	}

	head := p.renderHead(title)
	body := p.renderBody()

	return html.Doctype(
		html.HTML(
			html.Lang("en"),
			gomponents.Attr("data-theme", "dark"),
			html.Head(head...),
			html.Body(body...),
		),
	)
}
