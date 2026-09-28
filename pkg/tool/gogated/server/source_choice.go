package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func sourceChoice(
	value string,
	label string,
	selected string,
) gomponents.Node {
	input := []gomponents.Node{
		html.Type("radio"),
		html.Name(constant.SourceField),
		html.Value(value),
	}

	if value == selected {
		input = append(input, gomponents.Attr("checked", ""))
	}

	return html.Label(html.Input(input...), gomponents.Text(label))
}
