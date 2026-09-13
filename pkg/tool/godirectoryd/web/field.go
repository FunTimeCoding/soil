package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func field(
	name string,
	label string,
	kind string,
	required bool,
) gomponents.Node {
	input := []gomponents.Node{html.Type(kind), html.ID(name), html.Name(name)}

	if required {
		input = append(input, gomponents.Attr("required", ""))
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.Label(html.For(name), gomponents.Text(label)),
			html.Input(input...),
		},
	)
}
