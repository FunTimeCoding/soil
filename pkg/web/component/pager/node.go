package pager

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Pager) Node() gomponents.Node {
	if p.previousPath == "" && p.position == "" && p.nextPath == "" {
		return gomponents.Group(nil)
	}

	return html.P(
		html.Class(constant.PagerClass),
		gomponents.If(
			p.previousPath != "",
			html.A(html.Href(p.previousPath), gomponents.Text(p.previousLabel)),
		),
		gomponents.If(
			p.position != "",
			html.Span(
				html.Class(constant.PagerPositionClass),
				gomponents.Text(p.position),
			),
		),
		gomponents.If(
			p.nextPath != "",
			html.A(html.Href(p.nextPath), gomponents.Text(p.nextLabel)),
		),
	)
}
