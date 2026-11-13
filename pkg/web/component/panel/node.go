package panel

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Panel) Node() gomponents.Node {
	return html.Article(
		gomponents.If(p.identifier != "", html.ID(p.identifier)),
		gomponents.If(p.class != "", html.Class(p.class)),
		p.head(),
		gomponents.Group(p.content),
		gomponents.If(len(p.footer) > 0, html.Footer(p.footer...)),
	)
}
