package chip

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func (c *Chip) Node() gomponents.Node {
	classes := append([]string{constant.ChipClass}, c.kinds...)

	if c.toggle {
		classes = append(classes, constant.ChipToggleClass)
	}

	if c.active {
		classes = append(classes, constant.ChipActiveClass)
	}

	inner := []gomponents.Node{
		html.Class(join.Space(classes...)),
		gomponents.If(
			c.color != "",
			html.Style(join.Empty(constant.ChipColorVariable, ": ", c.color)),
		),
		gomponents.If(c.current, gomponents.Attr("aria-current", "true")),
		gomponents.Group(c.attributes),
		gomponents.Text(c.label),
		gomponents.Group(c.content),
	}

	switch {
	case c.toggle:
		return html.Button(
			append(
				[]gomponents.Node{
					html.Type("button"),
					gomponents.Attr(
						"aria-pressed",
						strconv.FormatBool(c.active),
					),
				},
				inner...,
			)...,
		)
	case c.link != "":
		return html.A(append([]gomponents.Node{html.Href(c.link)}, inner...)...)
	default:
		return html.Span(inner...)
	}
}
