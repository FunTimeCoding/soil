package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func dismissButtons(
	d *decision.Decision,
	target string,
) gomponents.Node {
	var buttons []gomponents.Node

	for _, v := range []constant.State{
		constant.StateDeclined,
		constant.StateIrrelevant,
		constant.StatePostponed,
	} {
		buttons = append(
			buttons,
			html.Button(
				gomponents.Attr(
					"hx-post",
					fmt.Sprintf(
						"%s?%s=%d&%s=%s",
						constant.DismissPath,
						constant.IdentifierParameter,
						d.Identifier,
						constant.StateParameter,
						v,
					),
				),
				gomponents.Attr("hx-target", target),
				gomponents.Attr("hx-swap", "outerHTML"),
				html.Class("secondary outline"),
				html.Style("margin-right: 0.5rem; width: auto;"),
				gomponents.Text(string(v)),
			),
		)
	}

	return html.P(gomponents.Group(buttons))
}
