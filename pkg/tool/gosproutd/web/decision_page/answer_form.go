package decision_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func answerForm(
	d *decision.Decision,
	alert string,
) gomponents.Node {
	return entry.New(
		actionPath(constant.AnswerPath, d.Identifier, constant.AnswerKindOther),
		"Answer",
	).
		Swap(target(d.Identifier), webConstant.SwapOuter).
		Inline().
		Alert(alert).
		Text(constant.ValueParameter, "Answer", "").
		Placeholder("something else, or a constraint").
		Alternative(
			actionPath(
				constant.AnswerPath,
				d.Identifier,
				constant.AnswerKindConstraint,
			),
			"You decide, within this",
		).
		Node()
}
