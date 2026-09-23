package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func answerForm(
	d *decision.Decision,
	target string,
) gomponents.Node {
	return html.Form(
		gomponents.Attr(
			"hx-post",
			fmt.Sprintf(
				"%s?%s=%d",
				constant.AnswerPath,
				constant.IdentifierParameter,
				d.Identifier,
			),
		),
		gomponents.Attr("hx-target", target),
		gomponents.Attr("hx-swap", "outerHTML"),
		html.Input(
			html.Type("text"),
			html.Name(constant.ValueParameter),
			gomponents.Attr("placeholder", "something else, or a constraint"),
		),
		html.Button(
			html.Type("submit"),
			html.Name(constant.KindParameter),
			html.Value(string(constant.AnswerKindOther)),
			html.Style("margin-right: 0.5rem; width: auto;"),
			gomponents.Text("Answer"),
		),
		html.Button(
			html.Type("submit"),
			html.Name(constant.KindParameter),
			html.Value(string(constant.AnswerKindConstraint)),
			html.Class("secondary"),
			html.Style("width: auto;"),
			gomponents.Text("You decide, within this"),
		),
	)
}
