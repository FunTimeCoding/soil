package card

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (c *Card) head() gomponents.Node {
	name := gomponents.Text(c.name)

	if c.link != "" {
		name = html.A(
			html.Class(constant.CardLinkClass),
			html.Href(c.link),
			gomponents.Text(c.name),
		)
	}

	title := html.H4(name)

	if c.size == constant.CardHeroClass {
		title = html.H1(name)
	}

	return html.Header(
		html.Class(constant.CardHeadClass),
		c.marker,
		title,
		gomponents.Group(c.trailing),
	)
}
