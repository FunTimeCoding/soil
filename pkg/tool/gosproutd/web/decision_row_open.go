package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func decisionRowOpen(d *decision.Decision) gomponents.Node {
	target := join.Empty("#", decisionIdentity(d.Identifier))

	return html.Tr(
		html.ID(decisionIdentity(d.Identifier)),
		html.Td(
			gomponents.Attr("colspan", "4"),
			html.Style("padding: 1.25rem;"),
			html.P(html.Strong(gomponents.Text(d.Question))),
			html.P(
				html.Small(
					html.Style("color: var(--pico-muted-color);"),
					gomponents.Text(
						join.Empty("If you never answer: ", d.DefaultAction),
					),
				),
			),
			choiceButtons(d, target),
			answerForm(d, target),
			dismissButtons(d, target),
			replyThread(d, target),
			html.P(
				html.Small(
					html.A(
						gomponents.Attr(
							"hx-get",
							fmt.Sprintf(
								"%s?%s=%d",
								constant.DecisionPath,
								constant.IdentifierParameter,
								d.Identifier,
							),
						),
						gomponents.Attr("hx-target", target),
						gomponents.Attr("hx-swap", "outerHTML"),
						html.Style("cursor: pointer;"),
						gomponents.Text("close"),
					),
				),
			),
		),
	)
}
