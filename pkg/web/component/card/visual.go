package card

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (c *Card) visual() gomponents.Node {
	switch {
	case c.portrait != "":
		return html.Div(
			html.Class(constant.CardVisualClass),
			html.Img(
				html.Src(c.portrait),
				html.Alt(c.name),
				html.Loading("lazy"),
			),
		)
	case c.emblem != "":
		return html.Div(
			html.Class(
				join.Space(constant.CardVisualClass, constant.CardFallbackClass),
			),
			gomponents.Text(c.emblem),
		)
	default:
		return nil
	}
}
