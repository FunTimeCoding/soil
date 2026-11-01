package card

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (c *Card) Node() gomponents.Node {
	classes := []string{constant.CardClass}

	for _, modifier := range []string{c.treatment, c.size} {
		if modifier != "" {
			classes = append(classes, modifier)
		}
	}

	if c.task {
		classes = append(classes, constant.CardTaskClass)
	}

	if c.class != "" {
		classes = append(classes, c.class)
	}

	meta := make([]gomponents.Node, 0, len(c.meta))

	for _, line := range c.meta {
		meta = append(
			meta,
			html.P(html.Class(constant.CardMetaClass), gomponents.Text(line)),
		)
	}

	return html.Article(
		gomponents.If(c.identifier != "", html.ID(c.identifier)),
		gomponents.If(c.preserve, extended.Preserve()),
		html.Class(join.Space(classes...)),
		gomponents.If(
			c.accent != "",
			html.Style(join.Empty(constant.CardAccentVariable, ": ", c.accent)),
		),
		gomponents.If(
			c.terms != "",
			gomponents.Attr(constant.SearchAttribute, c.terms),
		),
		c.visual(),
		html.Div(
			html.Class(constant.CardBodyClass),
			c.head(),
			gomponents.If(
				c.subtitle != "",
				html.P(
					html.Class(constant.CardSubtitleClass),
					gomponents.Text(c.subtitle),
				),
			),
			gomponents.Group(meta),
			gomponents.Group(c.state),
			gomponents.If(
				len(c.footer) > 0,
				html.Footer(
					html.Class(constant.CardFooterClass),
					gomponents.Group(c.footer),
				),
			),
		),
	)
}
