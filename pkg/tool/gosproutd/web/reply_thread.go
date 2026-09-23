package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func replyThread(
	d *decision.Decision,
	target string,
) gomponents.Node {
	var turns []gomponents.Node

	for _, v := range d.Turns {
		turns = append(
			turns,
			html.P(
				html.Small(
					html.Strong(
						gomponents.Text(join.Empty(string(v.Author), ": ")),
					),
					gomponents.Text(v.Content),
				),
			),
		)
	}

	return html.Div(
		gomponents.Group(turns),
		html.Form(
			gomponents.Attr(
				"hx-post",
				fmt.Sprintf(
					"%s?%s=%d",
					constant.ReplyPath,
					constant.IdentifierParameter,
					d.Identifier,
				),
			),
			gomponents.Attr("hx-target", target),
			gomponents.Attr("hx-swap", "outerHTML"),
			html.Input(
				html.Type("text"),
				html.Name(constant.ContentParameter),
				gomponents.Attr("placeholder", "reply without answering"),
			),
		),
	)
}
