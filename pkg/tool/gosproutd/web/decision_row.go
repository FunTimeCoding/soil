package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func decisionRow(d *decision.Decision) gomponents.Node {
	target := join.Empty("#", decisionIdentity(d.Identifier))
	var lines []gomponents.Node
	lines = append(
		lines,
		html.A(
			gomponents.Attr(
				"hx-get",
				fmt.Sprintf(
					"%s?%s=%d&%s=1",
					constant.DecisionPath,
					constant.IdentifierParameter,
					d.Identifier,
					constant.OpenParameter,
				),
			),
			gomponents.Attr("hx-target", target),
			gomponents.Attr("hx-swap", "outerHTML"),
			gomponents.Attr("role", "button"),
			html.Class("outline secondary"),
			html.Style("cursor: pointer; display: block; text-align: left;"),
			gomponents.Text(d.Question),
		),
	)

	if d.Answer != "" {
		attribution := "you: "

		if d.AnswerChannel == constant.AnswerChannelConversation {
			attribution = "you, in chat, as I heard it: "
		}

		lines = append(
			lines,
			html.P(
				html.Small(
					html.Style("color: var(--pico-muted-color);"),
					gomponents.Text(join.Empty(attribution, d.Answer)),
				),
			),
		)
	}

	if d.ClearLine != "" {
		lines = append(
			lines,
			html.P(
				html.Small(
					gomponents.Text(join.Empty("✓ taken as: ", d.ClearLine)),
				),
			),
		)
	}

	return html.Tr(
		html.ID(decisionIdentity(d.Identifier)),
		html.Td(gomponents.Group(lines)),
		html.Td(
			html.Small(
				html.Style("color: var(--pico-muted-color);"),
				gomponents.Text(frameNames(d)),
			),
		),
		html.Td(html.Small(gomponents.Text(decisionStanding(d)))),
		layout.TimeCell(d.CreatedAt),
	)
}
