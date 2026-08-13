package palette

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func resultItem(r Result) gomponents.Node {
	var children []gomponents.Node
	children = append(
		children,
		html.Span(
			html.Class("palette-label"),
			highlightLabel(r.Command.Label, r.Positions),
		),
	)

	if r.Command.Description != "" {
		children = append(
			children,
			html.Span(
				html.Class("palette-description"),
				gomponents.Text(r.Command.Description),
			),
		)
	}

	if r.Command.Category != "" {
		children = append(
			children,
			html.Span(
				html.Class("palette-category"),
				gomponents.Text(r.Command.Category),
			),
		)
	}

	var attrs []gomponents.Node
	attrs = append(
		attrs,
		html.Class("palette-item"),
		gomponents.Attr("data-palette-item", ""),
	)

	if r.Command.SwapTarget != "" {
		attrs = append(
			attrs,
			gomponents.Attr("href", "#"),
			gomponents.Attr("hx-get", r.Command.Path),
			gomponents.Attr("hx-target", r.Command.SwapTarget),
			gomponents.Attr("hx-swap", "innerHTML"),
		)
	} else {
		attrs = append(attrs, gomponents.Attr("href", r.Command.Path))
	}

	attrs = append(attrs, gomponents.Group(children))

	return html.A(attrs...)
}
