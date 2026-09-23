package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func choiceButtons(
	d *decision.Decision,
	target string,
) gomponents.Node {
	var buttons []gomponents.Node

	for _, v := range d.Choices {
		buttons = append(
			buttons,
			html.Button(
				gomponents.Attr(
					"hx-post",
					fmt.Sprintf(
						"%s?%s=%d&%s=%s",
						constant.AnswerPath,
						constant.IdentifierParameter,
						d.Identifier,
						constant.ValueParameter,
						url.QueryEscape(v.Label),
					),
				),
				gomponents.Attr("hx-target", target),
				gomponents.Attr("hx-swap", "outerHTML"),
				html.Style("margin-right: 0.5rem; width: auto;"),
				gomponents.Text(v.Label),
			),
		)
	}

	if buttons == nil {
		return gomponents.Text("")
	}

	return html.P(gomponents.Group(buttons))
}
